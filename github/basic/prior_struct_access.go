package main

import "fmt"

type Human1 struct {
	name string
	age int
	phone string
}
type Employee struct {
	Human1

	speciality1 string

	phone string
}
func main() {
	Bob := Employee{Human1{"Bab",25,"13733208110"},"Python","13091120193"}
	fmt.Println("Bod's work phone is:",Bob.phone)

	fmt.Println("Bob's personal phone is:",Bob.Human1.phone)
}



