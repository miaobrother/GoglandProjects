package main

import "fmt"

func main() {//如果有很多调用defer，那么defer是采用后进先出模式
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}//9 8 7 6 5 4 3 2 1 0
}