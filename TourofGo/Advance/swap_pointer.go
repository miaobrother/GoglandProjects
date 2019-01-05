package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("In the swap: a = %d, b = %d\n", *a, *b)

}

func main() {

	a, b := 100, 200
	swap(&a, &b)
	fmt.Printf("In the main: a = %d,b = %d\n", a, b)

}
