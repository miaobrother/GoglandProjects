package main

import (
	"time"
	"fmt"
)

func main() {
	timerOne := time.NewTimer(time.Second * 2)

	<- timerOne.C
	fmt.Println("Timer One expired")


	timerTwo := time.NewTimer(time.Second)
	//fmt.Println(timerTwo)

	go func() {
		<- timerTwo.C
		fmt.Println("Timer two expired")
	}()

	stopTwo := timerTwo.Stop()
	if stopTwo{
		fmt.Println("Timer two stopped")
	}
}