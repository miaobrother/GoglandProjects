package main

import (
	"fmt"
)

type PeopleA struct {
	Name string
	Age  int
}

type StuA struct {
	Score int
	PeopleA
}

func main() {
	var sss = StuA{78, PeopleA{"chen", 1900}}
	fmt.Printf("The sss is %v\n", sss)
}
