package main

import "fmt"

//普通相加函数
func Add(a,b int) int  {
	return a + b
}

//给某一个类型绑定一个函数

type long int

func (tmp long) AddTwo(other int) long  {
	return tmp + long(other)
}
func main(){
	//普通调用
	var result int
	result = Add(10,20)
	fmt.Printf("The result is %d\n",result)

	//定义一个变量
	var tmp long = 2
	//调用方法格式： 变量名.函数（参数）
	re := tmp.AddTwo(3)
	fmt.Printf("The result is %d\n",re)
}