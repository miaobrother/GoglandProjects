package main

import "fmt"

func main() {
	counter := 0
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			counter += i //3 6 9 = 18
		} else if i%5 == 0 {
			counter += i
			fmt.Printf("wo shi di yige :%d\n", counter)
		}
	}
	fmt.Println(counter)
}
