package main

import (
	"fmt"
	"reflect"
)

type UserOne struct {
	Id int
	Name string
}

type Manager struct {
	UserOne
	title string
}

func main() {
	m := Manager{UserOne{1,"zhangsan"},"lisi"}
	t := reflect.TypeOf(m)
	//fmt.Printf("%#v\n",t.Field(1))
	fmt.Printf("%#v\n",t.FieldByIndex([]int{0,1}))
}

