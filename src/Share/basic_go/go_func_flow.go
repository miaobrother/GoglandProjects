package main

import (
	"fmt"
)

func testOne(i int) int {
	if i == 1{
		return 1
	}
	return i + testOne(i - 1)
}

func main()  {
	var sum  int
	sum = testOne(100)
	fmt.Println("sum is",sum)
	fmt.Println("main")
}
