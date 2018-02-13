package main

import "fmt"

func add(a *int) int  {
	*a = *a +1
	return *a
}

func main() {
	x := 3
	fmt.Println("x=  ", x)

	x1 := add(&x)

	fmt.Println("x+1 = ",x1)

	fmt.Println("x = ",x)
}
