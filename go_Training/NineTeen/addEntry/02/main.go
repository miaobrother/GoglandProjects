package main

import "fmt"

func main() {
	var Mapp = map[int]string{
		1 :"zhangsan",
		2 : "nihao",
	}
	delete(Mapp,3)
	//Mapp = make(map[int]string)
	fmt.Println(Mapp)
}