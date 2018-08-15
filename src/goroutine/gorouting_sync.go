package main

import (
	"fmt"
	"time"
)

func Hell(exi chan bool) {
	fmt.Println("hello world")
	exi <- true
}

func main() {
	var exitChan chan bool

	exitChan = make(chan bool)
	//fmt.Println(exitChan)
	go Hell(exitChan)

	fmt.Println("main thread terminal")
	time.Sleep(time.Second)

}
