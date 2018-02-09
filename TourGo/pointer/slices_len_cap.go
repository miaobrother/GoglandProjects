package main

import "fmt"

func main() {

	s := []int{2,3,4,5,6,7,8}
	printSlice(s)

	s = s[:4]
	//fmt.Println(s)
	printSlice(s)

	s = s[2:]
	printSlice(s)
	//fmt.Println(s)
}

func printSlice(s [] int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}