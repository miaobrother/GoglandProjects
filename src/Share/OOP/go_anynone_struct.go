package main

import "fmt"

type PersonT struct {
	name string //名字
	sex byte //
	age int

}

type StudentT struct {
	PersonT  //只有类型 继承了Person的东西 匿名字段
	id int
	addr string

}
func main() {
	//顺序初始化
	var St StudentT = StudentT{PersonT{name:"zhangsan",sex:'m',age:19},1,"Handan"}
	fmt.Printf("The type is %#v\n",St)

}
