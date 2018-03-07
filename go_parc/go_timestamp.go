package main

import (
	"fmt"
	"time"
)

func main() {
	//使用Unix和UnixNano来分别获取从Unix起时shijian
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Printf("The Now is:%v\n",now)
	fmt.Printf("The Secs is:%v\n",secs)
	fmt.Printf("The Nanos is:%v\n",nanos)

	millis := nanos / 1e6
	fmt.Printf("The mis is:%v\n",millis)


	fmt.Println(time.Unix(secs,0))
	fmt.Println(time.Unix(0,nanos))
}