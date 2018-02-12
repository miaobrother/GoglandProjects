package main

import "fmt"

var s string = "hello"

func main() {
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)

	fmt.Printf("%s\n",s2)
}




