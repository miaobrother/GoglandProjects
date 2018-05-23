package main

import (
	"fmt"
	"time"
)

func write(ch chan int)  {
	for i := 0; i < 10;i ++{
		ch <- i
		fmt.Println("put data:",i)
	}
}

func read(ch chan int)  {
	for{
		var b int
		b = <- ch
		fmt.Println("i am a :",b)
		time.Sleep(time.Second)
	}
}

func main() {

	intChan := make(chan int,5)
	go write(intChan)
	go read(intChan)

	time.Sleep(10 * time.Second)
}