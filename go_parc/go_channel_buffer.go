package main

import "fmt"

func main() {
	messages := make(chan string,5)//带缓冲区的通道；

	messages <- "Zhangsan"
	messages <- "lis"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}