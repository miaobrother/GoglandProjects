package main

import (
	"fmt"
	//"time"
	"runtime"
)

func main() {

	go func() {
		runtime.GOMAXPROCS()
		for i := 0; i < 2; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ { //这是主的main 只打印两次hello

		runtime.Gosched()
		fmt.Println("hello")
		//time.Sleep(time.Second)
	}
}
