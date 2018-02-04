package main

import "fmt"

func main() {
	var x [10] int

	fmt.Println(x)//[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(x))
	fmt.Println(x[9])
	x[9] = 787
	fmt.Println(x[9])
}