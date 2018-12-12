package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is", time.Saturday)
	case today + 1:
		fmt.Println("Tomorrow is", time.Saturday)
	case today + 2:
		fmt.Println("In two days", time.Saturday)
	default:
		fmt.Println("Too far away.")

	}
}
