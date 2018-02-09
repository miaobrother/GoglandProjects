package main

import "fmt"

func main() {
	//var a [3]string
	a := [3]string{"1","2","3"}
	a[0] = "Name"
	a[1] = "Age"
	a[2] = "Sex"

	//fmt.Println(a[0],a[1],a[2])

	fmt.Println(a)


	primes := [6]int{2,3,4,5,6,7}
	fmt.Println(primes)
	//fmt.Printf("The primes type is:%v\n",primes)
}