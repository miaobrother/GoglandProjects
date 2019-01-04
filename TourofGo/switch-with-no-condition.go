package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("t.hour:", t.Hour())
	fmt.Printf("The t is %v\n", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening.")
	}
}
