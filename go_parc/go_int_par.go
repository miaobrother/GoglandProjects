package main

import (
	"strconv"
	"fmt"
)

func main() {
	//使用ParseFloat解析浮点数，64表示说明使用64位精度来表示

	f,_ := strconv.ParseFloat("1.234",64)
	fmt.Println(f)

	//对于ParsInt函数，0 表述从字符串推断整型进制
	i,_ := strconv.ParseInt("123",0,64)
	fmt.Println(i)

	k,_ := strconv.Atoi("124")
	fmt.Println(k)

}