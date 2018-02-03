package main

import "fmt"




var s = []int{1,2,4,4,5,6}
func main() {

	for i,v := range s{
		fmt.Printf("2**%d = %d\n",i,v)
	}

}
