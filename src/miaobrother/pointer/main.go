package main

import (
	"fmt"
)

func testPointer()  {
	var a int = 200
	var b *int = &a
	fmt.Printf("the b point value is: %v\n",*b)  //200
	*b = 1000
	fmt.Println(&a,b)
}

func modify(a *int)  {
	*a = 65
}
func testPointerOne()  {
	var b int = 10
	p := &b
	fmt.Printf("The p value before is:%v\n",*p)//10
	modify(p)
	fmt.Println(b)
}
func main() {
	//var a int32
	//a = 256
	//fmt.Println("Tha a address is:",&a,a)
	//
	//var b *int32
	//b = &a
	//fmt.Println(&b,&a)
	//testPointer()
	testPointerOne()
}
