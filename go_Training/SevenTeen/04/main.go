package main

import "fmt"

func main() {
	var x [256]byte

	for i:=0;i <len(x); i ++{
		x[i] = byte(i)
	}
	for i,v :=range x{
		fmt.Printf("%v -%T -%b\n",v,v,v,)
		if i > 50{
			break
		}
	}
}