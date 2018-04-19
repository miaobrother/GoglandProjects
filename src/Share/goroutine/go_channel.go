package main

import (
	"fmt"
	"time"
)
var ch = make(chan int)

func processOne()  {
	Printer("hello")
	ch <- 777

}

func Printer(str string)  {
	for _,data := range str{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")


}

func processTwo()  {
	<- ch
	Printer("world")

}
func main() {
	go processOne()
	go processTwo()

	for{

	}
}