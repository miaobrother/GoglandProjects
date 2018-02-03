package main

import "fmt"

type Int int


func Add(a,b int) int  {
	return a+b
}

func (i *Int)Add(a,b int) {
	*i = Int(a + b)
	return
}

func main()  {
	//c := Add(100,200)
	//fmt.Println(c)

	var a Int
	a.Add(100,200)
	fmt.Println(a)
}