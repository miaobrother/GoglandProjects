package main

import (
	"fmt"
	//"math"
	"sync"
	"time"
	"math/rand"
)

var wg sync.WaitGroup

var counter int

func main() {
	wg.Add(2)

	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()

	fmt.Println("Final Counter:",counter)
}

func incrementor(s string)  {
	rand.Seed(time.Now().UnixNano())//改变源数值

	for  i := 0; i < 20 ; i ++{
		x := counter
		x ++
		time.Sleep(time.Duration(rand.Intn(3) )* time.Millisecond)// 随机睡眠 0秒至 3秒
		counter += x
		fmt.Println(s,i,"Counter:",counter)
	}
	wg.Done()
}