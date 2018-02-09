package main

import "fmt"


type Var1 struct {
	x int
	y int
}

func main() {
	v := Var1{1,2}
	fmt.Println(v.x)
	v.x = 4
	fmt.Println(v.x)
}