package main

import (
	"fmt"
	//"os"
	//"unicode"
)

type point struct {
	x,y int
}

func main() {
	p := point{1,2}
	//下面的%v打印了一个point结构体的对象的值
	fmt.Printf("%v\n",p)


	// 如果所格式化的值是一个结构体对象，那么`%+v`的格式化输出
	// 将包括结构体的成员名称和值
	fmt.Printf("%+v\n",p)

	// 格式化输出将输出一个值的Go语法表示方式
	fmt.Printf("%#v\n",p)

	//使用 %T 来输出一个值的数据类型
	fmt.Printf("%T\n",p)

	fmt.Printf("%t\n",true)

	fmt.Printf("%d\n",123)

	fmt.Printf("%2f\n",78.90)


	 fmt.Printf("%e\n",12340000.0)
}