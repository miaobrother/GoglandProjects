package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//Go信号通知通过向一个channel发送 os.Signal来实现
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)

	signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)

	go func() {
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)

		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}