package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	//这个语法创建一个新的结构体变量
	fmt.Println(person{"zhangsan",18})

	//使用成员及值的方式来初始化结构体变量
	fmt.Println(person{name:"Chenmiao",age:17})

	//未显示赋值的成员初始值为0值
	fmt.Println(person{name:"lisi"})

	//可以使用&来获取结构体变量的地址
	fmt.Println(&person{name:"nimei",age:19})

	s := person{name:"Sean",age:50}
	fmt.Println(s.name)


	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)


}