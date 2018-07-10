package main

import "fmt"

type Student struct {
	id   int
	sex  byte
	age  int
	name string
}

func main() {
	var stu Student = Student{1, 'n', 19, "zhangsan"}
	fmt.Println(stu)
	s1 := Student{age: 10, name: "zhangs"}
	fmt.Println(s1)

	var p *Student = &Student{2, 'm', 89, "miaozi"}
	fmt.Println(p)

}
