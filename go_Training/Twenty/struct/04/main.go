package main

import (
	"fmt"
	//"crypto/x509/pkix"
)

type person struct {
	Frist string
	Last string
	Age int
}

type doubleZero struct {
	person
	Frist string
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person :person{
			Frist:"Zhang",
			Last:"YaLei",
			Age:24,
		},
		Frist : "Double",
		LicenseToKill:true,
	}

	p2 := doubleZero{
		person:person{
			Frist: "Chen",
			Last:"MiaoMiao",
			Age:19,
		},
		Frist : "If looks could kill",
		LicenseToKill:false,
	}

	fmt.Println(p1.Frist,p1.Last,p1.Age,p1.LicenseToKill)
	fmt.Println(p2.Frist,p2.Last,p2.Age,p2.LicenseToKill)

}