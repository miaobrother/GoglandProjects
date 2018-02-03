package main

import "fmt"

type People struct {
	Age  int
	Name string
}

type Student struct {
	Age  int
	Name string
	Next *People
}

func testlist() {
	var s Student
	s.Age = 100
	s.Name = "zhang"
	s.Next = &People{
		Age:  100,
		Name: "chen",
	}
	fmt.Printf("s:%+v\n", s)
	fmt.Printf("next:%v\n", *(s.Next))
}

func main() {
	testlist()
}
