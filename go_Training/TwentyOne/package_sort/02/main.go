package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zhang","Li","Chen","zhao"}
	fmt.Println(s)

	sort.Sort(sort.StringSlice(s))

	fmt.Println(s)
}