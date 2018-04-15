package main

import "fmt"

func main() {
	var id [50]int

	for i := 0;i < len(id);i ++{
		id[i] = i +1
		fmt.Printf("the i is %d, the v is %d\n",i,id[i])
	}
	/*
	for i,v := range id{
		fmt.Printf("the i is %d, the v is %d\n",i,v)
	}
	*/
}