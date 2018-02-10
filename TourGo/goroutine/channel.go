package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将 sum 送入c
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(len(s))
	c := make(chan int)
	go sum(s[:5], c)
	x := <-c //从c接受并赋值给x
	go sum(s[5:], c)
	y := <-c

	fmt.Println(x, y, x+y)
}
