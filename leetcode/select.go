package main

import "fmt"

func main() {
	var ch1,ch2,ch3 chan int

	var i2 int
	select {
	case i1 := <- ch1:
		fmt.Printf("received",i1,"from ch1\n")
	case ch2 <- i2:
		fmt.Printf("sent",i2,"to ch2\n")
	case i3,ok := (<-ch3):
		if ok{
			fmt.Printf("received",i3,"from ch3")
		}else {
			fmt.Printf("ch3 is closed\n")
		}
	default:
		fmt.Printf("no commuication\n")
	}
}