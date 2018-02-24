package main

import (
	"fmt"
	//"go/constant"
)

var (
	nums   = []int{2, 7, 11, 15}
	target = 9
	res    = []int{}
)

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println("This is the res:",nums[i], nums[j])
				//return res
			}
		}
	}
	return res
}

func main() {
	twoSum(nums, 9)
}
