package main

import "fmt"

type Person struct {
	name string
	age int
	sex byte
}

type Student struct {
	*Person //指针类型
	id int
	addr string
}

func main()  {
	sTwo := Student{&Person{"Jack",19,'f'},19,"handan"}

	fmt.Println(sTwo.name,sTwo.addr,sTwo.age,sTwo.sex)

	//先定义一个变量

	var sThreee Student
	sThreee.Person = new(Person)
	sThreee.name = "zhangsan"
	sThreee.sex ='f'
	sThreee.age = 19
	fmt.Println(*sThreee.Person)
}