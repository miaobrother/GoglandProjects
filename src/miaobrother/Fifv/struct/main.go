package main

import (
	//"go/types"
	"fmt"
)

type Int int
type Test struct {
	A int
	b int
}

type Student struct {
	Age   int
	Name  string
	Sex   string
	Grade string
	Sroce int
	a Int
	b Test
}

func testInt() {
	var a Int
	var b int
	a = Int(b)
	fmt.Println(a, b)
}

func testStruct() {
	var s Student
	s.Age = 100
	s.Name = "Zhang"
	s.Sroce = 80
	s.Sex = "man"
	s.b.A = 100


	fmt.Printf("name:%s age:%d sroce:%d sex:%s\n", s.Name, s.Age, s.Sroce, s.Sex)//name:Zhang age:100 sroce:80 sex:man
	fmt.Printf("%+v\n",s)//{Age:100 Name:Zhang Sex:man Grade: Sroce:80}
}

func main() {
	testStruct()
}
