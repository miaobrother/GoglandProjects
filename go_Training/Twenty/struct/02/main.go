package main

import "fmt"


type person struct {
	first string
	last string
	age int
}

func main() {
	p1 := person{"James", "Bond",20}

	p2 := person{"Miss","Moneypenny",19}

	fmt.Println(p1.first,p2.last,p1.age)
}