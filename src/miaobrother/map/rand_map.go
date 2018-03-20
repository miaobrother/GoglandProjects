package main

import (
	"fmt"
	"time"
	"math/rand"
	//"expvar"
)

func main() {
	rand.Seed(time.Now().UnixNano())//生成随机种子

	var m map[string] int = make(map[string]int,1024)//初始化一个Map

	for i :=0;i < 256; i++{
		key := fmt.Sprintf("stud%d",i)
		value := rand.Intn(256)
		m[key] = value
	}
	for k,v := range m{
		fmt.Printf("The key is:%v,The value is:%v\n",k,v)
	}
	fmt.Println(len(m))

}