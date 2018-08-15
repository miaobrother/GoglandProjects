package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) Print() {
	fmt.Println(s)
}

func (s Student) Set(name string, age int, score float32) {
	s.Age = age
	s.Name = name
	s.Score = score
}

func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	fmt.Println(val, kd)
	fmt.Println("kd is ", kd) //struct
	fmt.Println("reflect struct is ", reflect.Struct)

	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取字段数量

	fields := val.NumField()
	fmt.Println("The fields is ", fields)

	//获取字段类型

	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("%d %v\n", i, val.Field(i).Kind())
	}

	//获取方法数量

	numMethod := val.NumMethod()
	fmt.Println("Method num is ", numMethod)

	//反射调用Print方法
	var params []reflect.Value

	val.Method(0).Call(params)
}

func main() {
	var s Student = Student{
		Name:  "zhangsan",
		Age:   19,
		Score: 93.8,
	}

	TestStruct(s)

}
