package main

import "fmt"

type Verext struct {
	X int
	Y int
}

var (
	v1 = Verext{1, 2}
	v2 = Verext{X: 1}
	v3 = Verext{}
	p  = &Verext{1, 7}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
