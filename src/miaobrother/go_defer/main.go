package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}

func (p *person) fullName()  {
	fmt.Printf("%v %v\n",p.firstName,p.lastName)
}

func main() {
	p := &person {
		firstName:"zhangsan",
		lastName:"lisi",
	}
	defer p.fullName()
	fmt.Println("Welcome  ")
}
