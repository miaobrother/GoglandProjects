package main

import (
	"fmt"
)

func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}
func main() {
	testFor()
}
