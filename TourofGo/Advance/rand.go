package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randValue() [10]int {
	rand.Seed(time.Now().Unix())
	var dict [10]int
	for i := 0; i < len(dict); i++ {
		dict[i] = rand.Intn(200)
	}
	return dict
}

//数组做函数参数，默认是值传递
func sortBool(dict [10]int) [10]int {
	for i := 0; i < len(dict); i++ {
		for j := 0; j < len(dict)-i-1; j++ {
			if dict[j] > dict[j+1] {
				dict[j], dict[j+1] = dict[j+1], dict[j]
			}
		}
	}
	return dict
}

func main() {
	//randValue()
	fmt.Printf("The sortedbool is:%v\n", sortBool(randValue()))

}
