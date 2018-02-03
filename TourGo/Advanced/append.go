package main

import "fmt"

func main() {
	var s[]int
	//printSlice1(s)
	//
	//s = append(s,0)
	//printSlice1(s)

	s = append(s,1,2,3,4,5,6)
	printSlice1(s)
	
}

func printSlice1(s [] int)  {
	fmt.Printf("len = %d ,cap = %d %v\n",len(s),cap(s),s)
}