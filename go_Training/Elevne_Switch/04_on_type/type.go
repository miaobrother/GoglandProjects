package main

import (
	"fmt"
	//"go/ast"
)

type contact struct {
	greeting string
	name string
}

func SwitchOnType(x interface{})  {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Mical")
	var t = contact{"Good to se you,","Zhangsan"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}

