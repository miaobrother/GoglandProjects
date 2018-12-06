package main

import (
	"fmt"
)

var (
	Tobe   bool   = false
	MaxInt uint64 = 1<<64 - 1
)

func main() {
	fmt.Printf("Type: %T value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T value is: %v\n", MaxInt, MaxInt)
}
