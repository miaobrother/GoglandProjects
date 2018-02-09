package main

import (
	"fmt"
	"math"
	//"encoding/binary"
)

func Pow1(x,y, lim float64) float64  {
	if v:= math.Pow(x,y); v < lim{
		return v
	}else {
		fmt.Printf("%g >= %g\n",v,lim)

	}
	return lim
}

func main() {
	fmt.Println(
		Pow1(2,2,10),
		Pow1(10,3,20),
	)
}