package main

import "fmt"

func main() {
	rating := map[string]float64{"c":3.5,"Go":4.5,"Python":4.5,"c++":3}// 第一种方法。

	seaarch,ok := rating["c"]
	if ok != true{
		fmt.Println("nil")
	}else {
		fmt.Println(seaarch)
	}

	fmt.Println(rating)

	rating2 := make(map[string]float64)//第二种方法。
	rating2["css"] = 4
	fmt.Println(rating2)
}
