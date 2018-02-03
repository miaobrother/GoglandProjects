package main

import "fmt"

func main() {
	pipe := make(chan int, 3)
	//define a channel and int last three

	pipe <- 10
	pipe <- 20
	pipe <- 30
	pipe <- 40
	var a int
	a = <-pipe

	fmt.Printf("%d\n", a)

}
