package main

import (
	"fmt"
	"math/rand"
	"time"
	//"go/ast"
)

func NewTest() chan int {
	c := make(chan int, 1)

	rand.Seed(time.Now().UnixNano())

	go func() {
		time.Sleep(time.Second)
		c <- rand.Intn(10)
	}()
	return c
}
func rands() int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)
	//fmt.Println(r)
	return r

}
func main() {

	//fmt.Println(<-t)
	for {
		r := rands()
		//fmt.Println(r)

		t := NewTest()
		i := <-t
		if r != i {
			fmt.Println("r bu等于 t", r, i)
		} else {
			fmt.Println("deng yu ", r, i)
			break
		}

	}

}
