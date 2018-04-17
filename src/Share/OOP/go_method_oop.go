package main

import "fmt"

type person struct {
	name string
	sex byte
}

func (tmp *person) PrintfInfo()  {//这是一个方法
	fmt.Printf("这是person的，The name is :%s sex is :%c\n",tmp.name,tmp.sex)
}



func (tmp *student) PrintfInfo()  {//这是一个方法
	fmt.Printf("这是student的，The name is :%s sex is :%c\n",tmp.name,tmp.sex)
}

type student struct {
	person //匿名字段
	id int
}

func main() {
	s := student{person{"zhangsan",'f'},20}
	s.PrintfInfo() //The name is :zhangsan sex is :f 可以看到student继承了person的类 实现了PrintInfo方法的调用
}
