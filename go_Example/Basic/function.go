package main

import (
	"fmt"
)

func plus(a, b int) int {
	return a + b
}

func threePlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2= ", res)
	threeres := threePlus(1, 2, 3)
	fmt.Println(threeres)
}