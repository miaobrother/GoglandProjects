package main

import (
	"fmt"
	//"log"
	"math/rand"
	"time"
)

func CreatNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(1e4)
		if num >= 1000 {
			break
		}
	}
	*p = num
}

func GetNum(s[]int,num int)  {
	s[0], s[1], s[2], s[3] = num/1000, (num%1000)/100, (num%100)/10, num%10

}
func InputNum(randX []int) {
	var num int
	keySlice := make([]int, 4)
	for {

		for {
			fmt.Println("请输入4位数: ")
			fmt.Scan(&num)
			if 999 < num && num <= 10000 {
				break
			}
			fmt.Println("输入数字不合法")
		}
		//fmt.Println("The num is ",num)

		GetNum(keySlice, num)
		//fmt.Println(keySlice)
		n := 0
		for i :=0; i < 4;i++{
			if keySlice[i] > randX[i]{
				fmt.Printf("第%d位大了一点\n",i+1)
			}else if keySlice[i] < randX[i] {
				fmt.Printf("第%d位小了一点\n",i+1)
			}else {
				fmt.Printf("第%d位数相等\n",i+1)
			}
		}
		if n == 4{
			fmt.Println("all right")
		}
	}
}
func main() {
	var randNum int
	CreatNum(&randNum) //创建一个四位数
	randSlice := make([]int,4)
	fmt.Println("The randNum is:",randNum)
	GetNum(randSlice,randNum)
	fmt.Println("randSlice is : ", randSlice)

	InputNum(randSlice)//获取用户输入的一个4位数


}
