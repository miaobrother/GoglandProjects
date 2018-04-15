package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子
	t := time.Now().UnixNano()
	fmt.Println(t)
	rand.Seed(t)//使种子值变化
	for i := 0;i <5;i ++{
		fmt.Println("the rand is :",rand.Intn(10))//限制在 10以内的数
		fmt.Println("the rand is :",rand.Int())//随机一个很大的数
	}

}