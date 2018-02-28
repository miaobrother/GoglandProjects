package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm Zhang,Hi"

	fmt.Printf("The Position of \"Zhang\" is:")

	fmt.Printf("%d\n",strings.Index(str,"Zhang"))

	fmt.Println(strings.Replace(str,"Hi","iH" , -1))
}