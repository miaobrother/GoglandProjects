// fibonacci number 1 1 2 3 5 8 13
package main

import (
	"fmt"
)

func fibonacci(ch chan <- int,q <-chan bool)  { //ch 是只写  q 是只读
	x,y := 1,1
	for{
		//监听channel中数据的流动
		select {
		case ch <- x:
			x,y = y ,x+y
		case flag:= <- q:
			fmt.Println("flag = ",flag)
			return
		}
	}
}

func main() {
	ch := make(chan int) //实现数字间计算
	q := make(chan bool) //实现通知程序结束标示
	go func() {
		for i := 0;i < 10;i++{
			num := <- ch
			fmt.Println(num)
		}
		q <- true
	}()
	fibonacci(ch, q)
}