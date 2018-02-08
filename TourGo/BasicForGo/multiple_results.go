package main

import "fmt"

func swap(x,y string) (string,string)  {//定义一个返回两个字符串的函数。
	return y,x
}

func main() {
	a,b := swap("Hello","World")//声明并初始化swap函数。
	fmt.Println(a,b)  //res is :World Hello
}
