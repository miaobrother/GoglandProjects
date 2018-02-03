package main

import "fmt"

func main() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	close(queue)
	fmt.Println(len(queue))

	for elem := range queue {
		fmt.Println(elem)
	}
}
