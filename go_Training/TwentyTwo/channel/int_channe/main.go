package main

import (
	"fmt"
	//"go/ast"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0;i < 10; i ++{
			c <- i
		}
		close(c)
	}()

	for n := range c{
		fmt.Println(n)
	}
}