package main

import (
	"fmt"
	"time"
)


func main() {
	 t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Printf("Now is %v Good afternoon\n",t.Hour())
	default:
		fmt.Printf("evening")

	}
}
