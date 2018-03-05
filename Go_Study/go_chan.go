package main

import "fmt"

func main() {
	//range chan中的数据
	queue := make(chan string,4)

	queue <- "one"
	queue <- "two"
	queue <- "three"

	close(queue)
	for elem := range queue{
		fmt.Println(elem)
	}
}