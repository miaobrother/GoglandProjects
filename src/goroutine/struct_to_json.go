package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type UserFour struct {
	UserName string `json:"user_name"`
	SecondName string `json:"second_name"`
	Age int
	Birthday string
	Sex string
	Email string
	Phone string
}

/*
 结构体 到  json
 */
func toStruct()  {
	user := &UserFour{
		UserName:"chenmiao",
		SecondName:"miaozi",
		Age:24,
		Birthday:"1993-02-18",
		Sex:"女",
		Email:"7291@126.com",
		Phone:"13000000000",
	}

	data,err := json.Marshal(&user)
	if err != nil{
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n",string(data))
}

func main() {
	toStruct()
}