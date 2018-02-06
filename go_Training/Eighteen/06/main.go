package main

import "fmt"

func main() {
	customerNumber := make([]int,3)

	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 16

	fmt.Println(customerNumber[0],customerNumber[1])


	greeting := make([]string,3,5)

	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}