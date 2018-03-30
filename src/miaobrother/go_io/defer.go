package main

import (
	"fmt"
)

func funcA()int  {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func main() {
	fmt.Println(funcA())
}