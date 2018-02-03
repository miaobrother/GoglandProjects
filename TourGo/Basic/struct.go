package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Sex   string
	Grage string
	Score int
}

func testStruct() {
	var s Student
	s.Age = 26
	s.Name = "Jack"
	s.Score = 79
	s.Sex = "Man"
	s.Grage = "Hello"

	//fmt.Println(s)
	fmt.Printf("%+v", s)
	fmt.Printf("name:%s,age:%d,score:%d,sex:%s,grage:%s\n", s.Name, s.Age, s.Score, s.Sex, s.Grage)
}

func main() {
	testStruct()

}
