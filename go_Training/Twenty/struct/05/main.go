package main

import "fmt"


type person struct {
	Name string
	Age int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greet()  {

	fmt.Println("I'm just a regular")
}

func (dz doubleZero) Greet()  {
	fmt.Println("Miss Moneypenny")
}

func main() {
	p1 := person{
		Name:"I am a Fle",
		Age :8776,
	}

	p2 := doubleZero{
		person:person{
			Name:"james",
			Age:9888,
		},
		LicenseToKill:true,
	}
	p1.Greet()
	p2.Greet()
	p2.person.Greet()
}