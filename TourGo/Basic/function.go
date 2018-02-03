package main

import (
	"fmt"
)

func testSum(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(testSum(43, 12))
}
