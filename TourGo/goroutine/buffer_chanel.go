package main

import (
	"fmt"
	"time"
)
func main() {
	ch := make(chan int,2)//定义一个带缓冲区的channel 大小是2
	ch <- 1
	fmt.Println(<-ch)
	time.Sleep(3000 * time.Millisecond)
	ch <- 2
	fmt.Println(<-ch)
}
