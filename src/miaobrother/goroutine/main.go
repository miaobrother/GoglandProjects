package main

import (
	"fmt"
	"time"
)

func calc() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("run", i, "times")
	}
	fmt.Println("calc finish")
}

func main() {
	go calc() //此时是创建了一个goroutine

	fmt.Println("i exited..") // 程序首先打印这一输出
	time.Sleep(11 * time.Second)
}
