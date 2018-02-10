package main

import "fmt"

func fib() func() int {
	return
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
