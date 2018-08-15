package main

import "fmt"

type PersonT struct {
	name string //名字
	sex  byte   //
	age  int
}

type StudentT struct {
	PersonT //只有类型 继承了Person的东西 匿名字段
	id      int
	name    string //和Person同名了
	addr    string
}

func main() {
	/*
		//顺序初始化
		var St StudentT = StudentT{PersonT{name:"zhangsan",sex:'m',age:19},1,"Handan","hebei"}

		fmt.Printf("The type is %#v\n",St)
	*/
	//声明一个变量
	var s StudentT
	s.sex = 'f'
	s.age = 18
	s.name = "jack"
	s.PersonT.name = "lisi"
	s.addr = "hengshui"

	fmt.Printf("The s is %+v\n", s)
}
