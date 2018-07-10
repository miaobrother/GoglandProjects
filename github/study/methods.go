package main

import (
	"fmt"
	"math"
)

type VertexTwo struct {
	X, Y float64
}

func (v *VertexTwo) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &VertexTwo{3, 4}
	fmt.Println(v.Abs())

}
