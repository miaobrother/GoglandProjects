package main

import "fmt"

type VertexS struct {
	X int
	Y int
}

func main() {
	v := VertexS{1,2}
	v.X = 4
	fmt.Println(v)
}
