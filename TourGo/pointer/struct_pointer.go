package main

import "fmt"

type Vartext struct {
	X int
	Y int
}

func main() {
	v := Vartext{1,2}
	p := &v
	//fmt.Println((*p).X)//两个语句是一样的意思。
	//fmt.Println(p.X)
	p.X = 1e9
	fmt.Println(p.X)
}
