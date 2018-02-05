package main

import "fmt"

func main() {
	xs := []int{1,2,5,7,9}

	for i ,value :=range xs{
		fmt.Println(i,"-",value)
	}
}