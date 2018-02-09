package main

import (
	"fmt"
	"math"
)
//var ni string := hao
func main() {
	var x,y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("x is %v\n y is %v\n z is %v\n",x,y,z)

}


