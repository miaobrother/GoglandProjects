package main

import (
	"fmt"
	"go/ast"
)

func sum(a []int,c chan int)  {
	tatol:= 0

	for _,v := range a{
		tatol += v
	}
	c <- tatol
}

func main() {
	a := []int{7,2,8,-9,4,0}

	c := make(chan int)

	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)

	x,y := <-c, <-c
	fmt.Println(x,y,x+y)
}