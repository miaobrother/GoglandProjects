package main

import "fmt"

func bubble_sort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}

	}
}

func main() {
	a := []int{1, 4, 25, 3, 89, 67, 0}
	bubble_sort(a)
	fmt.Println(a)
}
