package main

import (
	"fmt"
	//"go/ast"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("go")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
