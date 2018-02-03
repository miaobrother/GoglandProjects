package main

import (
	"fmt"
	"miaobrother/Day1/calc"
)

func main() {
	var sum int
	var sub int
	sum = calc.Add(2, 3)
	sum, sub = calc.Calc(2, 3)
	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Printf("hello")

}
