package main

import "fmt"

func main() {
	a := make([]int,5)
	printMake("a",a)
	//fmt.Printf("a is :%v\n",a)
	 b := make([]int,0,5)
	 printMake("b",b)
	 //fmt.Printf("b is :%v\n",b)
	 c := b[:2]
	 printMake("c",c)
	//d := c[2:5]
	//printSlice("d", d)
}

func printMake(s string,x[]int)  {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}