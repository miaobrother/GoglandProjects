package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print()  {
	for i := 0; i < 2; i ++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
var ch = make(chan int)

func sell(ch chan int)  {
	for{
		num := <- ch
		fmt.Println("sell", num,"bread")
	}
}

func produce(p chan int)  {
	for{
		num := rand.Intn(10)
		t := time.Duration(time.Duration(num))
		fmt.Println("Product",num,"bread")
		ch <- num
		time.Sleep(time.Second * t)
	}
}

func main() {
	go print()
	var input string
	fmt.Scanln(&input)
	fmt.Println("End")

	go sell(ch)
	go produce(ch)

	var input2 string
	fmt.Scanln(&input2)
	fmt.Println("end")
}
