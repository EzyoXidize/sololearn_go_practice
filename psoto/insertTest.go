package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"bytes"
)

type Students struct {
	Id 		int 		`json:"id"`
	Name 	string		`json:"name"`
	Gender 	int			`json:"gender"`
	Class	int			`json:"class"`
}

type Studentslice struct {
	ReqStu []Students	`json:"reqstu"`
}

func main() {

	url := "http://127.0.0.1:443/insert"

	//stu := student{
	//	Id 			:	56,
	//	Name		:	"aaa",
	//	Gender 		:	0,
	//	Class 		:	11,
	//}
	var stu Studentslice
	stu.ReqStu = make([]Students,2)
	stu.ReqStu[0].Id = 215
	stu.ReqStu[0].Name="ayang"
	stu.ReqStu[0].Gender = 1
	stu.ReqStu[0].Class = 13
	stu.ReqStu[1].Id = 216
	stu.ReqStu[1].Name="alog"
	stu.ReqStu[1].Gender= 0
	stu.ReqStu[1].Class= 14




	bs, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json err:", err)
	}

	 //        fmt.Println(string(bs))
	 req := bytes.NewBuffer([]byte(bs))
	 //tmp := `{"id": 5,"name":"aaa","gender":"0","class":"11"}`
	 //req = bytes.NewBuffer([]byte(tmp))

	 client := &http.Client{}
	 request, err := http.NewRequest("POST", url, req)

	 if err != nil {
	 	fmt.Println(err.Error())
	 	return
	 }

	 request.Header.Set("Content-Type", "application/json;charset=UTF-8")



	 resp, err := client.Do(request)


	 body, err := ioutil.ReadAll(resp.Body)

	 if err != nil {
	 	fmt.Println("响应错误")
	 }

	 fmt.Println(string(body))

	 fmt.Println(err)


	/*
	client := &http.Client{}

	req := `{"id": 5,"name":"aaa","gender":"0","class":"11"}`

	req_new := bytes.NewBuffer([]byte(req))
	request, _ := http.NewRequest("POST", "local:7000/insert", req_new)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
	*/
}