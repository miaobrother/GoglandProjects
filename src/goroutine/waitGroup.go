package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1) //添加一个执行的线程
		go calc(&wg, i)
	}

	wg.Wait()
	fmt.Println("all goroutine done")
}

func calc(w *sync.WaitGroup, i int) {
	fmt.Println("calc", i)

	time.Sleep(time.Second)

	w.Done()
}
