package main

import (
	"fmt"
	//"go/types"
)

func main() {
	var x interface{}

	switch i:= x.(type) {
	case nil:
		fmt.Printf("x 的类型是 %T\n",i)
	case int:
		fmt.Printf("x 的类型是 %T\n",i)
	case float64:
		fmt.Printf("x  的类型是 %T\n",i)
	default:
		fmt.Printf("未知类型")
	}
}