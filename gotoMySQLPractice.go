package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"

	"log"

)


type request struct {
	requests []byte
}



const DB_driver  = "root:123456@tcp(localhost:3306)/testdb?charset=utf8"


func (r request ) findsql(db *sql.DB)  {
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		fmt.Println("error:", err)
	} else {
	}
	for rows.Next() {
		var id 		int
		var sname   string
		var gender 	int
		var class 	int
		err = rows.Scan(&id,&sname,&gender,&class)
		fmt.Printf("id:%d  name : %s  gender: %d  class : %d\n",id,sname,gender,class)
	}
	return
}


func (r request) insertsql(db *sql.DB,idnum int,sname string,gendernum int, classnum int)  {
	stmt,err := db.Prepare("INSERT INTO students(id,sname,gender,class) values(?,?,?,?)")
	if err != nil {
		fmt.Println("stmt error")
	}

	result,err := stmt.Exec(idnum,sname,gendernum,classnum)
	if err != nil {
		fmt.Println("exec error")

	}
	lastInsertId, err := result.LastInsertId()
	rowsAffected, err := result.RowsAffected()
	fmt.Println("lastInsertId", lastInsertId)
	fmt.Println("影响的行数：", rowsAffected)

}


func (r request) updatesql(db *sql.DB,classnum int,idnum int)  {
	stmt,err := db.Prepare("update students set class = ? WHERE students.id=?")

	if err != nil {
		fmt.Println("update error")
	}
	//changelow,err := stmt.RowsAffected()
	//if err != nil {
	//}
	stmt.Exec(classnum,idnum)
	//fmt.Println("共计",changelow,"行受到影响")

}

func (r request) deletesql(db *sql.DB,num int)  {
	stmt,err := db.Prepare("delete from students where students.id=?")
	if err != nil {
		fmt.Println("delete error")
	}
	stmt.Exec(num)

}

func main() {

	req := new(request)

	db, err := sql.Open("mysql", DB_driver)
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}


	//req.insertsql(db,142,"xiaoqq",0,14)
	//req.deletesql(db,141)
	// req.updatesql(db,classnum,idnum)

	req.findsql(db)

}


/*
func postrequest(postreq []byte) request {
	select {
	case reqRe,err := regexp.MustCompile( `(SELECT [*]+ FROM students)`):
		reqs := reqRe.Find(postreq)
		if err != nil {
			fmt.Println("req error")
		result := request{}
		result.requests = append(reqs)

		}

	case reqRe,err := regexp.MustCompile(``)

	}
}
*/

