package main

import (
	"fmt"
)

func main() {
	var m map[string] int
	//fmt.Println(m)
	if m == nil{
		m = make(map[string]int)
		fmt.Printf("m = %v\n",m)

		m["student01"] = 1
		m["student02"] = 2
		m["student03"] = 3
		fmt.Println("after",m)

		b := m
		b["student01"] = 1000
		fmt.Println(b)
		fmt.Printf("after change map is %v\n",m)

	}
}
