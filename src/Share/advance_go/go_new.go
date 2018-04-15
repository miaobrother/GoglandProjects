package main

import "fmt"

func main() {
	aa := 100

	var ppp *int
	ppp = &aa//这是只想一个合法的内存
	fmt.Printf("The ppp is %v\n",&ppp)
	fmt.Printf("The aa addr is %v\n",&aa)


	//ppp 是*int，指向int类型
	ppp = new(int)
}