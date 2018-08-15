package main

import (
	"fmt"
	"time"
)

type worker struct {
	in   chan int
	done chan bool
}

func doworker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		go func() {
			done <- true
		}()
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doworker(id, w.in, w.done)
	return w
}
func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
	time.Sleep(time.Second)

	//n := <-c
	//fmt.Println(n)
}

func bufferedChannel() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3

}
func channelColse() {
	return

}

func main() {
	chanDemo()
	bufferedChannel()
	channelColse()
}
