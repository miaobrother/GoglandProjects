package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc 3.1 3.45 zhangsan lisi 9.01 823.1"

	// 解释正则表达式
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil{
		fmt.Println("存在异常 请检查")
		return
	}
	res := reg.FindAllString(buf,-1)
	fmt.Printf("The result is %v\n",res)
}
