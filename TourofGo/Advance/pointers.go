package main

import "fmt"

func main() {

	var abc = 10
	abd := 10
	var p *int
	p = &abc

	var ppp *int
	ppp = new(int)
	*ppp = 18
	fmt.Printf("the ppp is %v\n", *ppp)
	fmt.Printf("The p is %d\n", *p)
	fmt.Printf("&abd = %p\n", &abd)
	fmt.Printf("abd = %d\n", abd)
	fmt.Printf("&abc = %p\n", &abc)
	fmt.Printf("abc = %d\n", abc)
	//变量的内存，变量的地址

	i, j := 42, 2701

	*p = 21
	fmt.Println(i)

	p = &j

	*p = *p / 37

	fmt.Println(j)
}
