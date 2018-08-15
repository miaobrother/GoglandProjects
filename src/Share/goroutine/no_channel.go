package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)

	fmt.Printf("len ch is %d,cap is %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程,%d is ", i)
			ch <- i
		}
	}()
	time.Sleep(time.Second * 2)

	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("Num = ", num)
	}
}
