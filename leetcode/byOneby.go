package main

import (
	"fmt"
	"time"
)

func pNum(s string, c chan int) {
	for _, i := range s {
		time.Sleep(time.Second)
		fmt.Println(string(i))
		c <- 1
	}

}
func pStr(s string, c chan int) {
	for n, i := range s {
		<-c
		fmt.Println(string(i))
		if n == len(s)-1 {
			c <- 1
			c <- 1

		}
	}
}
func main() {
	c := make(chan int, 2)
	go pNum("123456789", c)
	go pStr("abcdefghi", c)
	for {
		if len(c) == 2 {
			fmt.Println("Print is over.")
			return
		}
	}
}
