package main

import (
	"fmt"
	"math/rand"
	"time"
	//"go-ethereum/metrics"
)

func CreateNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}
	//fmt.Println("The num is", num)
	*p = num
}
func getEcheNum(s []int, num int) { //集成到一个函数中，求出一个四位数的个 十 百 千 位 数字
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
	fmt.Println(s)

}

func GetNumTerminal(randSlice []int) {
	{

		keySlice := make([]int, 4)
		var num int
		for {
			fmt.Printf("Please input a num: ")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
		}
		//fmt.Println("The num is :",num)
		getEcheNum(keySlice, num)
		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了一些\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了一点\n", i+1)
			} else {
				fmt.Printf("第%d猜对了\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("全部猜对了")

		}
	}
}

func main() {
	var randNum int
	CreateNum(&randNum)
	//fmt.Println("randNum:", randNum)

	randSlice := make([]int, 4)
	getEcheNum(randSlice, randNum)

	GetNumTerminal(randSlice)

	/*
		该逻辑实现的是拆分四位数
		nOne := randNum/1000
		nTwo := randNum%1000 /100
		nThree := randNum %100 / 10
		nFour := randNum %10
		fmt.Printf("The nOne is :%d\n",nOne)
		fmt.Printf("The nOne is :%d\n",nTwo)
		fmt.Printf("The nOne is :%d\n",nThree)
		fmt.Printf("The nOne is :%d\n",nFour)
	*/
}
