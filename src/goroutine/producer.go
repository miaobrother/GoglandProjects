package main

import (
	"fmt"
	//"golang.org/x/net/http2"
)

func produce(channel chan int)  {
	for i := 0;i < 10;i ++{
		channel <- i  //将i写入到channel中
	}
	close(channel)
}

func main() {
	ch := make(chan int)
	go produce(ch)

	for{
		v,ok := <-ch
		if ok == false{
			break
		}
		fmt.Println("Received",v,ok)
	}
}
