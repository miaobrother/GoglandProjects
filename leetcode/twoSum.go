package main

import (
	"fmt"
)

var (
	num   = []int{2, 7, 11, 15, 3, 6}
	target = 9
	res    = []int{1}
)

func twoSum(num []int, target int) []int {

	for i := 0; i < len(num); i++ {
		for j := 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				fmt.Println("This is the res:", num[i], num[j])
				//return res
			}
		}
	}
	return res
}

func main() {
	twoSum(num, 9)
}
