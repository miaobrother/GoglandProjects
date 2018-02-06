package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) fullName() string  {
	return p.first + p.last
}

func main() {
	p1 := person{"zhang","Yalei",20}

	p2 := person{"Chen","Miao",19}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}