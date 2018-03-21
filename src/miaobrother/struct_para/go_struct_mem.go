package main

import "fmt"

type Test struct {
	a int32
	b int32
	c int32
	d int32
}



func main() {
	var t Test
	fmt.Printf("The a addr is %p\n",&t.a)
	fmt.Printf("The a addr is %p\n",&t.b)
	fmt.Printf("The a addr is %p\n",&t.c)
	fmt.Printf("The a addr is %p\n",&t.d)

}