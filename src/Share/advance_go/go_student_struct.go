package main

import (
	"fmt"
	//"net/http"
)


type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main() {
	var stuOne Student = Student{//顺序初始化 每个程序必须初始化
		1,"zhangsan",'m',26,"handan",
	}
	fmt.Printf("The stuOne is %#v\n ",stuOne)

	// 指定成员初始化，没有初始化的成员，自动赋值
	sTwo := Student{addr:"hengshui"}

	fmt.Printf("The stuOne is %#v\n ",sTwo)




	var ptr *Student = &Student{//顺序初始化 每个程序必须初始化
		1,"lisi",'f',26,"handan",
	}
	fmt.Printf("The ptr is %#v\n ",*ptr)
}