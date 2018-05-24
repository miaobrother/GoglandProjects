package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type UserSix struct {
	UserName string `json:"user_name"`
	Age int
	Sex string
}

/*
map to  json
 */

func TMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})

	m["user_name"] = "zhangsan"
	m["age"] = 19
	m["sex"] = "男"
	data ,err := json.Marshal(m)
	if err != nil{
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n",data)
}

func main() {
	TMap()

}