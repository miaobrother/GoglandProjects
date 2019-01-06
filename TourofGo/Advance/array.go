package main

import "fmt"

func main() {
	var ids [50]int
	for i := 1; i < len(ids); i++ {
		ids[i] = i
		fmt.Printf("The i is %d The value is %d\n", i, ids[i])
	}
}
