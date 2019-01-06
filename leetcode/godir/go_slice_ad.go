package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	t := time.Now().UnixNano()
	rand.Seed(t)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}

}

func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	fmt.Println("排序后是:", s)
}
func main() {
	n := 10
	s := make([]int, n)
	InitData(s)
	fmt.Println("排序前是:", s)

	BubbleSort(s)
}
