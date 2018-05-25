package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	age string
}

func main() {
	t,err := template.ParseFiles("./index.html")
	if err != nil{
		fmt.Println("parse file err")
		return
	}

	p:= Person{
		Name:"zhangsan",
		age:"19",
	}
	if err := t.Execute(os.Stdout,p);err != nil{
		fmt.Println("There was an error:",err.Error())
	}
}
