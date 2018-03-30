package main

import "fmt"

type fileOne struct {
	name string
}
type pipe struct {
	name string
}

func main() {
	var f fileOne
	var p pipe
	fmt.Printf("the f type is %T\n",f)
	fmt.Printf("The p type is %T\n",p)
}
