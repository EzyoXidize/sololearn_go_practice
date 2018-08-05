package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

type Server struct {
}

type Person struct {
	id 		int 		`xorm:"not null pk autoincr int"`
	name 	string		`xorm:"not null VARCHAR(100)"`
}

var engine *xorm.Engine

func main()  {
	var err error
	http.HandleFunc("/",requestsHandler)

	err = http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
	
}


func requestsHandler(w http.ResponseWriter, r *http.Request)  {

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
	method, found3 := r.Form["Method"]
	name, found2 := r.Form["name"]
	id, found1 := r.Form["id"]


	result,_ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	fmt.Println("%s\n",result)


	bytes,_ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
	
	person := new(Person)

	person.findsql(engine,person)
	//person.insertsql(engine,person)
	//person.deletesql(engine,person)
	//person.updatesql(engine,person)

}


func (p *Person) findsql(engine *xorm.Engine,person *Person)  {

	result,err := engine.Where("id=?",1).Get(person)
	if err != nil {
		fmt.Println("find error")
	}
	fmt.Println(result)
}


func (p *Person) insertsql(engine *xorm.Engine,person *Person)  {

	affected,err := engine.Insert(person)
	if err != nil {
		fmt.Println("insert error")
	}
	fmt.Println(affected)
}


func (p *Person) deletesql(engine *xorm.Engine,person *Person)  {

	affected_delete , err := engine.Delete(person)
	if err != nil {
		fmt.Println("delete error")
	}
	fmt.Println(affected_delete)
}


func (p *Person) updatesql(engine *xorm.Engine,person *Person)  {


	affected_update,err := engine.Id(1).Update(person)
	if err != nil {
		fmt.Println("update error")
	}
	fmt.Println(affected_update)
}
