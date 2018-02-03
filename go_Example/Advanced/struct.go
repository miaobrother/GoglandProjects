package main

import (
	"fmt"
	//"encoding/xml"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 19})

	fmt.Println(person{name: "Frand"})

	fmt.Println(&person{name: "Ann", age: 20})
	s := person{name: "scan", age: 90}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 91
	fmt.Println(sp.age)
}
