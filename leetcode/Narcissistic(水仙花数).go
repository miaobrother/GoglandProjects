package main

import (
	"fmt"
)

func main() {
	CalcShui()

}

func ShuiXianHua(n int) bool {
	one, two, three := n/100, (n/10)%10, n%10
	//fmt.Println(one ^ 3 + two ^ 3 + three ^ 3)
	if one*one*one+two*two*two+three*three*three == n {
		return true

	}
	return false

}

func CalcShui() {
	for i := 100; i < 1000; i++ {
		if ShuiXianHua(i) == true {
			fmt.Printf("%d is a narc\n", i)

		}
	}

}
