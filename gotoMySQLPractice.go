package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)


type Person struct {
	Id 		int 		`xorm:"not null pk autoincr int"`
	Name 	string		`xorm:"not null VARCHAR(100)"`
}


func main()  {
	var err error
	http.HandleFunc("/",requestsHandler)

	err = http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}

}


func requestsHandler(w http.ResponseWriter, r *http.Request)  {

	var engine *xorm.Engine
	engine, err := xorm.NewEngine("mysql", "root:123456@/xorm?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return
	}

	if err := engine.Ping() ; err != nil {

		fmt.Println(err)

		return
	}

	engine.ShowSQL(true)
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(5)


	r.ParseForm()

	result,_ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	fmt.Println("%s\n" , result)

	var person  Person
	json.Unmarshal(result, &person)

	person.findsql(engine,person)
	//person.insertsql(engine,person)
	//person.deletesql(engine,person)
	//person.updatesql(engine,person)

	w.Write([]byte("well done"))

}


func (p *Person) findsql(engine *xorm.Engine,person Person)  {

	result,err := engine.Where("id=?",person.Id).Get(person)
	if err != nil {
		fmt.Println("find error")
	}
	fmt.Println(result)
}


func (p *Person) insertsql(engine *xorm.Engine,person Person)  {

	affected,err := engine.Insert(person)
	if err != nil {
		fmt.Println("insert error")
	}
	fmt.Println(affected)
}


func (p *Person) deletesql(engine *xorm.Engine,person Person)  {

	affectedDelete , err := engine.Where("id=?",person.Id).Delete(person)
	if err != nil {
		fmt.Println("delete error")
	}
	fmt.Println(affectedDelete)
}


func (p *Person) updatesql(engine *xorm.Engine,person Person)  {


	affectedUpdate,err := engine.Where("id=?",person.Id).Update(person)
	if err != nil {
		fmt.Println("update error")
	}
	fmt.Println(affectedUpdate)
}
