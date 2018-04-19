package main

import (
	"fmt"
	"time"
	//"go-ethereum/common/number"
	//"go/constant"
	//"golang.org/x/net/html/atom"
)
func main() {
	chOne := make(chan int)
	quit := make(chan bool)

	go func() {
		for{
			select {
			case numOne := <-chOne:
				fmt.Println("num =",numOne)
			case <- time.After(time.Second *3):
				fmt.Println("超时了")
				quit <- true
			}
		}
	}()
	for i := 0;i < 5;i ++{
		chOne <- i
	}
	<-quit
	fmt.Println("程序结束")
}