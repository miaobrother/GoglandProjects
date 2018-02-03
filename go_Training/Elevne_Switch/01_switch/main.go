package main

import "fmt"

func main() {
	switch "Mhi" {
	case "Danil":
		fmt.Println("Wassup Daniel")
	case "Lisi":
		fmt.Println("Wassup Lisi")
	case "Zhangsan":
		fmt.Println("WASSUP  Zhangsan")
	default:
		fmt.Println("Have you no friends?")

	}
}
