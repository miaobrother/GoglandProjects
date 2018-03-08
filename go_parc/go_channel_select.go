package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msgOne := <-c1:
			fmt.Println("received", msgOne)
		case msgTwo := <-c2:
			fmt.Println("received", msgTwo)
		}
	}
}
