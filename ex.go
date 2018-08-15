package main

import "fmt"

var s = []int{1, 2, 3, 4, 5, 6, 7}

func For(s []int) (sum int) {
	for i := 0; i < len(s); i++ {
		sum = sum + s[i]
	}
	return sum
}

func recursion(x []int) (y int) {
	if len(x) == 0 {
		return 0
	} else {
		return x[0] + recursion(x[1:])

	}
	return

}

func main() {
	fmt.Println(For(s))
	fmt.Println(recursion(s))

}
