package main

import (
	"fmt"
	"errors"
)

var b = "hello"//两种方式是一样的
func main() {
	m := `hello
world`
	//b := "hello"
	b = "c" + b[1:]
	fmt.Println(m)
	fmt.Printf("%s\n",b)


	err := errors.New("Please input you num:\n")
	if err != nil{
		fmt.Print(err)
	}
}
