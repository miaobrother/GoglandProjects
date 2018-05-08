package main

import "fmt"

func main() {
	var mapOne map[int]string
	fmt.Println(len(mapOne))


	mapTwo := make(map[int]string)
	fmt.Println(len(mapTwo))
}