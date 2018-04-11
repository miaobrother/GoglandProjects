package main

import (
	"fmt"
	"time"
)

var  o chan int


func ready(w string, sec int)  {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w,"is ready")
	o <- 1
}




var cOne = make(chan int,10)

var oo string
func p()  {
	oo = "Hello World"
	cOne <- 0

}

var a2 string

var cTwo = make(chan int)

func f2()  {
	a2 = "hello my world"
	<-cTwo
}

func main() {
	o = make(chan int)
	go ready("Tee",2)
	go ready("Coffee",1)
	fmt.Println("i am waiting ,but not too long")

	//从channel中获取数据
	<- o
	<- o
	go p()
	<- cOne
	fmt.Println(oo)

	go f2()
	cTwo <- 0
	fmt.Println(a2)
}
