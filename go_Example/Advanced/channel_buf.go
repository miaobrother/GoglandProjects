package main

import "fmt"

func main() {
	messages := make(chan string, 5)
	messages <- "buffered"
	messages <- "channel"
	messages <- "three"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
