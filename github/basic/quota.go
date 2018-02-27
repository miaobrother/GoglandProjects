package main

import (
	"fmt"
)

func main() {
	//a := []int{0,0,0,} //define a slice
	//fmt.Println(a)
	//
	//b := make([]int,6)
	//fmt.Println(b)
	//c := new([]int)
	//
	//fmt.Println(&c)
	s := "abc"
	btos := []byte(s)
	btos[0] = 'A'
	fmt.Println(s)
}