package main

import "fmt"

func main() {

	adc := []int{1,2,3,4,5,6,7}
	sOne := adc[3:]
	fmt.Printf("The sOne is %v\n ",cap(sOne))
}