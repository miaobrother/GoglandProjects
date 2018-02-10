package main

import "fmt"

func fibb(n int,c chan int)  {
	x,y := 0,1
	for i :=0; i <n; i++{
		c <- x
		x,y = y , x+y
	}
	close(c)
}

func main() {
	c := make(chan int,10)//生成一个cap为10 的channel
	go fibb(cap(c),c)
	for i:= range c{
		fmt.Println(i)
	}
}