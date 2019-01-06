package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	var sorts [10]int

	for i := 0; i < len(sorts); i++ {
		sorts[i] = rand.Intn(100)
		fmt.Printf("%d,", sorts[i]) //随机产生了10个随机正数
	}
	fmt.Printf("\n")
	//冒泡排序 相邻的两个元素进行升序比较

	for i := 0; i < len(sorts)-1; i++ {
		for j := 0; j < len(sorts)-1-i; j++ {
			if sorts[j] > sorts[j+1] {
				sorts[j], sorts[j+1] = sorts[j+1], sorts[j]
			}
		}
	}
	fmt.Println("排序后的数据")
	for i := 0; i < len(sorts); i++ {
		fmt.Printf("%d,", sorts[i])
	}
	fmt.Printf("\n")
}
