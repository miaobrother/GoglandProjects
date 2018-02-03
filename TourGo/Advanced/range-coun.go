package main

import (
	"fmt"
)

func main() {
	pow := make([]int ,10)//define the slices

	for i:= range pow {//for loop the slices
		pow[i] = 1 << uint(i)
	}

	for _,v := range pow{
		fmt.Printf("%d\n",v)
	}

}