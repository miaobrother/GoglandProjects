package main

import (
	"fmt"
)

// 4,9,3,7,11,1
func bullble_sort(s []int) {
	for i := 0; i < len(s)-1; i++ { //趟数
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}

	}
}
func select_sort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := i; j < len(s); j++ {
			if s[j] > s[i] {
				s[j], s[i] = s[i], s[j]
			}
		}
	}
}

func main() {
	s := []int{4, 9, 3, 7, 11, 1}
	bullble_sort(s)
	//select_sort(s)
	fmt.Println(s)
}
