package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)
	messages := make(chan string)

	go func() {messages <- input}()

	msg := <- messages
	fmt.Println(msg)
}