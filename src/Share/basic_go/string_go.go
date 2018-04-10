package main

import (
	"fmt"
	//"strconv"
)

var str string//声明字符串
var strOne = "i am a boy"
var strTwo = "i am a man"


func main() {
	str = "test"
	ch := str[0]
	l := len(str)
	fmt.Println(str, string(ch),l)
	fmt.Println()//空一行

	my := "my name is 张亚磊"
	fmt.Println(len(my))//20

	for i := 0;i < len(my);i ++{
		fmt.Println(string(my[i]))
	}
	fmt.Println()
	for i,s := range my{
		fmt.Println(i,"Unicode(",s,")string-",string(s))
	}


	r := []rune(my)//该类型是一个切片， rune是一种数据类型，是int32

	fmt.Println("rune=",string(r))

	for i := 0; i < len(r);i++{
		fmt.Println("r[",i,"] = ",r[i],"string = ",string(r[i]))
	}
	fmt.Println(strOne,strTwo)
}
