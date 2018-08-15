package main

import "fmt"

type Student struct {
	Age  int
	Name string
	Next *People
}
type People struct {
	Age  int
	Name string
}

func testlist() {
	var s = Student{19, "haoze", &People{20, "miaoze"}}
	fmt.Printf("The s is %#v\n", *s.Next)
	var ss Student
	ss.Name = "zhangsan"
	ss.Age = 10
	ss.Next = &People{
		10,
		"lisi",
	}
}
func main() {
	//fmt.Printf("The Student is %#v\n",Student{})
	//fmt.Printf("The People is %#v\n",People{})
	testlist()
}
