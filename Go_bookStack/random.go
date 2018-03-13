package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i:= 0;i <3;i ++{
		a := rand.Int()
		fmt.Printf("%d\n",a)
	}
	for i := 0;i <5 ;i++{
		r := rand.Intn(8)
		fmt.Printf("%d\n",r)//所生成的是小于 8的整数
	}
	fmt.Println()

	times := int64(time.Now().Nanosecond())
	rand.Seed(times)

	for i := 0;i <10; i++{
		fmt.Printf("%0.2f\n",100*rand.Float32())//显示小于1的float类型
	}
}
//5577006791947779410
//18446744073709551616