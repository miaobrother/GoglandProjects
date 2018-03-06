package main

import (
	"fmt"
	"time"
)

func main() {
	jobes := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobes
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}

		}
	}()

	for j := 1; j <= 6; j++ {
		time.Sleep(time.Nanosecond * 1000000000)
		jobes <- j
		fmt.Println("sent job", j)
	}
	close(jobes)
	fmt.Println("sent all jobs")

	<-done
}
