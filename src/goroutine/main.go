package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	ch <- 1000
}

func customer(ch chan int) {
	data := <-ch
	fmt.Printf("The customer data is %v\n", data)
}
func main() {
	var ch chan int
	ch = make(chan int) //不带缓冲区的 得先去chan中取数据
	go producer(ch)
	go customer(ch)
	ch <- 100
	fmt.Printf("Type of ch is %v\n", ch)
	time.Sleep(time.Second)
}
