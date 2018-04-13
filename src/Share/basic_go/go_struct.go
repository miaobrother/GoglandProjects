package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age int32
}

type course []string

type student struct {
	name string
	age int32
	grade string
}

type class struct {
	student
	course
	int32
	teacher string
}
type goods struct {
	name string "goods_name"
	price float64 "goods_price"
}

func main() {
	//var s []string

	stu02 := class{student{"xiaolei",18,"大学"},[],27,"zhangsan"}
	stu02.course  = make([]string,6)
	stu02.course[0] = "math"
	stu02.course = append(stu02.course,"china")
	stu02.int32 = 7
	fmt.Println(stu02)
	//这是单独声明和赋值的
	var p1 person
	p1.name = "zhangyalei"
	p1.age = 26

	p2 := person{"xiaohao",20}//直接声明和赋值的

	p3 := person{age:27,name:"zhangsan"}//通过key value 值初始化，顺序随意
	fmt.Println(p1,p2,p3)






}





