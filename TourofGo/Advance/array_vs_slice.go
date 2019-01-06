package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[1:3:5]
	fmt.Println("a is ", a)
	fmt.Printf("The s len is %d\n The s cap is %d\n", len(s), cap(s))
}
