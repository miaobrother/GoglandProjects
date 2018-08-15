package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Stu struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

func main() {
	var stu = &Stu{
		"hello",
		16,
		97.1,
	}
	data, err := json.Marshal(&stu)
	if err != nil {
		log.Fatal("Marshal is failed..")
		return
	}
	fmt.Printf("json data is %s\n", data)
	//fmt.Printf("The data is %+v\n",string(data))
	var s1 Stu
	err = json.Unmarshal(data, &s1)
	if err != nil {
		fmt.Printf("The unmarshal failed ,err:%v\n", err)
		return
	}
	fmt.Printf("s1:%#v\n", s1)
}
