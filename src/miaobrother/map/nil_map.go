package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Printf("The map is %v\n", m)
	if m == nil {
		m = make(map[string]int, 16) //初始化Map
		fmt.Printf("Tha a is %v\n", m)
	}
}
