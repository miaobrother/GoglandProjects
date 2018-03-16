package main

import (
	"fmt"
)

func isNarcissistic(n int) bool  {
	First := n % 10 //个位数
	Second := (n /10) % 10 //十位数
	Third := n / 100
	if First*First*First + Second*Second*Second +Third*Third*Third == n{
		return true
	}

	//fmt.Printf("n:%d f is %d,t is %d,three is %d\n",n,First,Second,Third)
	return false
}

func expaleTwo()  {
	for i := 0;i < 1000;i++{
		if isNarcissistic(i) == true{
			fmt.Printf("%d is Narcissistic\n",i)
		}
	}
}
func main() {
	expaleTwo()

}