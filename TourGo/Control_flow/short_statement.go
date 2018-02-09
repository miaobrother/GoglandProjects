package main

import (
	"fmt"
	"math"
)

func Pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v <= lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(Pow(2, 3, 4),
		Pow(1, 4, 5),
	)
}
