package main

import "fmt"

type AB interface {
}

func main() {
	var a AB
	var b string
	a = b
	fmt.Println(a)
}
