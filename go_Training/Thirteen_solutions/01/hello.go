package main

import "fmt"

//func main() {
//	var name string
//
//	fmt.Print("Please input your want to be name:")
//	fmt.Scan(&name)
//
//	fmt.Println("Hello",name)
//}

func main() {
	var numOne int
	var numTwo int

	fmt.Print("Please enter a large number:")

	fmt.Scan(&numOne)

	fmt.Print("Please enter smaller number:")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "%", numTwo, " = ", numOne, numTwo)
}
