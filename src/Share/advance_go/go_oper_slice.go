package main

import "fmt"

func main() {

	slice := []int{1,2,3,4,5,6,7,8,9,0}
	s1 := slice[3:5:9]
	fmt.Printf("The s1 is %v\n",s1)
	fmt.Printf("The s1 cpa is %v\n",cap(s1))
}