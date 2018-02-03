package main

import (
	"fmt"
)

func main() {
	a :=make([] int ,5)
	fmt.Println("a is:",a)

	b := make([] int,0,5)
	fmt.Println("b is :",b)

	c := a[:2]
	fmt.Println("C is :",c)

	d := c[2:5]
	fmt.Println("D is:",d)

}