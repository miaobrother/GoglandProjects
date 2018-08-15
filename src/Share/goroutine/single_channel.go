package main

import (
	//"database/sql"
	"fmt"
)

func producer(out chan<- int) { //只写channel
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)

}

func consumer(in <-chan int) { //只读channel
	for num := range in {
		fmt.Println("num = ", num)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch) //主协程 先执行
}
