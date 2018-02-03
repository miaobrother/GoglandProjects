package main

import "fmt"

func binary_test() {
	var num int
	var num1 int
	var num2 int
	num = 1 << 10
	num1 = 1 << 20
	num2 = 1 << 30
	fmt.Printf("%v\n%v\n%v\n", num ,num1,num2)
}