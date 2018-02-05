package main

import "fmt"

func main() {
	myMap := map[string]string {
		"name":"zhangsan",
		"Age":"19",
		"Sex": "M",
	}
	fmt.Println(myMap)

	if val,exists := myMap["name"]; exists{
		fmt.Println("That value exists.")
		fmt.Println("val",val)
		fmt.Println("exists:",exists)
	}else {
		fmt.Println("That doesn'ot exists")
		fmt.Println("val",val)
		fmt.Println("exists:",exists)
	}

	fmt.Println(myMap)
}