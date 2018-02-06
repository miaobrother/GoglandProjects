package main

import "fmt"

func main() {
	data := []float64{1,2,3,4,5,6,7,8,9}
	n := average(data)
	fmt.Println(n)
	fmt.Printf("%T\n",n)
}

func average(sd []float64) float64  {
	tatol := 0.0
	for _,v := range sd{
		tatol += v
	}
	return tatol/float64(len(sd))
}