package main

import (
	"fmt"
	"reflect"
)

type UserTwo struct {
	Id int
	Name string
}

func (u UserTwo) HelloWorld(name string)  {
	fmt.Println("Hello ",name,"my name is ",u.Name)
}

func main() {
	u := UserTwo{1,"chenmiao"}
	//u.HelloWorld("jack")
	v := reflect.ValueOf(u)
	mv := v.MethodByName("HelloWorld")
	args := []reflect.Value{reflect.ValueOf("jok")}
	mv.Call(args)
	//fmt.Printf("The v is %#v\n",v)
}
