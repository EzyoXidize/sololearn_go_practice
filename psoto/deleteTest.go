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
	Name 	string		`json:"name,omitempty"`
	Gender 	int			`json:"gender,omitempty"`
	Class	int			`json:"class,omitempty"`
}

func main() {

	url := "http://127.0.0.1:443/delete"

	var stu = Students{
		Id:127,
	}


	bs, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json err:", err)
	}

	req := bytes.NewBuffer([]byte(bs))

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


}