package main

import "fmt"

func main() {

	var s1 = [5]int{1,2,3,4,5}
	var s2 = [5]int{1,2,3,4,5}

	if s1 == s2 {
		fmt.Println("The s1 and s2 equal")
	}else {
		fmt.Println("The s1 and s2 not equal")
	}
}
