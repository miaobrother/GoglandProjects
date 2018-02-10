package main

import (
	"math"
	"fmt"
)

type Vsex struct {
	x, y float64
}

func Abs(v Vsex)float64  {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vsex{3,4}
	fmt.Println(Abs(v))
}
