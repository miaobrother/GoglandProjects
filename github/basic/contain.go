package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "I am a chinaese,Please input your file:"
	fmt.Printf("T/F ? Does the string \"%s\" have prefix %s?\n",str,"Pl")
	fmt.Printf("%t\n",strings.HasPrefix(str,"Pl"))
	fmt.Println(strings.Contains(str,"Pl"))
}