package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int {9,4,7,1,3,8,34,0,-9}
	fmt.Println(n)

	//sort.Sort(sort.IntSlice(n))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println(n)
}