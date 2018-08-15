package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumArray(a [10]int) int {
	var sum int = 0
	/*
		for i := 0; i < len(a); i++{
			sum = sum +a[i]
		}
		return sum
	*/

	for _, val := range a {
		sum = sum + val
	}
	return sum
}

func testSumArray() {
	rand.Seed(time.Now().Unix())
	var gc [10]int
	for i := 0; i < len(gc); i++ {
		gc[i] = rand.Intn(1000)
	}
	sum := sumArray(gc)
	fmt.Printf("sum is %d\n", sum)
}
func main() {
	testSumArray()
}
