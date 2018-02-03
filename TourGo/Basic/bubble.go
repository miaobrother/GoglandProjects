package main

import (
	"fmt"
	"sort"
	//"container/ring"
)

func bubble_sort(a []int) { //冒泡排序
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func basic_sort(a []int) { //程序自带冒泡排序
	sort.Ints(a)
}

func select_sort(a []int) { //选择排序
	//var index int = i
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

func insert_sort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
}

func partion(a []int, left, right int) int {
	var i = left
	var j = right
	for i < j {
		for j > i && a[j] > a[left] {
			j--
		}
		a[j], a[left] = a[i], a[j]

		for i < j && a[i] < a[left] {
			i++
		}
		a[left], a[i] = a[i], a[left]
	}
	return i
}

func fastsort(a []int, left, right int) {
	if left >= right {
		return
	}
	mid := partion(a, left, right)

	fastsort(a, left, mid-1)
	fastsort(a, mid+1, right)
}

func main() {
	a := []int{38, 1, 5, 3, 9, 8}
	//bubble_sort(a)
	//basic_sort(a)
	//select_sort(a)
	//insert_sort(a)
	fastsort(a, 0, len(a)-1)
	fmt.Println(a)

}
