package main

import (
	"fmt"
	"math"
)

type Vecs struct {
	x , y float64
}

func (v Vecs) Abs() float64  {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vecs{3,4}
	fmt.Println(v.Abs())
}