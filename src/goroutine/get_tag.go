package main

import (
	"fmt"
	"reflect"
)

type Teacher struct {
	Name  string `json:"teach_name"`
	Age   int
	Score float32
}

func TeStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	tag := typ.Elem().Field(0).Tag.Get("json") //获取tag 只能根据TyprOf获取
	fmt.Printf("Tag:%s\n", tag)
}

func main() {
	var t Teacher = Teacher{
		Name: "zhangsan",
		Age:  19,
	}

	TeStruct(&t)
}
