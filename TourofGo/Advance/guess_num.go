package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int)  {
	rand.Seed(time.Now().Unix())

	var num int
	for {
		num = rand.Intn(10000)
		if num > 1000{
			break
		}
	}
	*p = num
}

func GetPeerNum(randSlice []int,randNum int)(ran []int)  {

	randSlice[0] = randNum / 1000
	randSlice[1] = randNum % 1000 /100
	randSlice[2] = randNum % 100 /10
	randSlice[3] = randNum % 100 % 10
	return ran
	//fmt.Printf("The randSlice is %v\n",randSlice)

}

func GameNum(newSlice []int) ( s []int){
	var inputNum int
	for{
		fmt.Printf("请输入你心目中的数字:")
		fmt.Scan(&inputNum)
		if 999 < inputNum && inputNum < 1e4{
			break
		}
		fmt.Println("不符合要求 请重新输入:")
	}
	GetPeerNum(newSlice,inputNum)
	fmt.Println("The newSlice is:",newSlice)
	//fmt.Printf("The inputNum is %d\n",inputNum)
	return newSlice
}

func main() {
	var randNum int
	randSlice := make([]int,4)
	newSlice := make([]int,4)

	CreateNum(&randNum)
	GetPeerNum(randSlice,randNum)
	GameNum(newSlice)
	fmt.Printf("The randNum is:%d\n",randSlice)
	n := 0
	for i := 0;i < 4;i ++{
		if randSlice[i] > newSlice[i]{
			fmt.Printf("第%d位大了一点\n",i+1)
		}else if randSlice[i] < newSlice[i]{
			fmt.Printf("第%d位小了一点\n",i+1)
		}else {
			fmt.Printf("guess right...")
			n++
		}
		if n == 4{
			fmt.Printf("all guess right")
			break
		}
	}
}