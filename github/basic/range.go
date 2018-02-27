package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3,4}
	for i,v := range a{
		if i == 0{
			a[1],a[2] = 999,999
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
}