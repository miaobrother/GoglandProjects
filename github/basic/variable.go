package main

import (
	"fmt"
)

func test()  {
	y := "Zhangsan"
	fmt.Println(&y)
}
func main() {
	s := "abc"
	fmt.Println(&s)
	s,y := "hello",20
	fmt.Println(&s,&y)
	test()
	fmt.Println(s)
}