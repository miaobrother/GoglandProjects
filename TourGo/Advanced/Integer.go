package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool  {
	return a < b
}//定义新的类型Integer

func main() {
	var a Integer = 1

	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
}
