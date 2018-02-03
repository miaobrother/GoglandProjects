package main

import (
	"fmt"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible 4")
	}

	if num := 12; num < 0 {
		fmt.Println(num, "is negatuve")
	} else if num < 10 {
		fmt.Println(num, "Has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
