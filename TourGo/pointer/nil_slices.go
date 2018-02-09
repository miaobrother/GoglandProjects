package main

import "fmt"

func main() {
	var b []int
	fmt.Printf("The b is :%v The len is:%v, The cap is:%v\n",b,len(b),cap(b))
	if b == nil{
		fmt.Println("The b is nil ")
	}
}
