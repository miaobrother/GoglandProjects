package main

import "fmt"

const (
	a = iota
	b = 0
	c

)
const (
	d = iota
	e
	f

)
const (
	a1 = 1 << iota
	a2 = 1 << iota
	a3 = 1 << iota
	a4 = 1 << iota
	)

func main()  {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

}