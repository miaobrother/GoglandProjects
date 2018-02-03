package main

import (
	"time"
	"fmt"
	//"go/ast"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("tick at ", t)
		}
	}()
	time.Sleep(time.Millisecond * 2000)
	ticker.Stop()
	fmt.Println("Ticker is Stopped")
}
