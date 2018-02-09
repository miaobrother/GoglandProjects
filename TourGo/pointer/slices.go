package main

import "fmt"

func main() {
	primes := [10] int{1,2,3,4,5,6,7,8,9,0,}

	var s []int = primes[1:5]
	fmt.Println(s)
}

