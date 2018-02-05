package main

import "fmt"

func main() {
	var myGreeting map[string]string
	myGreeting = make(map[string] string)

	fmt.Println(len(myGreeting))

	myGreeting["name"] = "zhangsan"
	myGreeting["age"] = "19"

	fmt.Println(&myGreeting)

	//fmt.Println(myGreeting == nil)
}