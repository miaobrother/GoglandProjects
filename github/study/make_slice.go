package main

import "fmt"



func main() {
	c := make([]string,10)
	fmt.Printf("The c is %v\n ",c)
	a := make([]int,10)//
	fmt.Println(a)
	b := make([]int,0)
	fmt.Println(b)
	fmt.Printf("The b type is %d\n",len(b))
}
