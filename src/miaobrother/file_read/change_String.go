package main

import (
	"fmt"
)

func main() {
	var s string = "abcdef"
	for i,v := range s{
		fmt.Printf("%d,%c\n",i,v)
	}

	var strToByte []byte
	strToByte = []byte(s)
	strToByte[0] = 'i'//[97 98 99 100 101 102]
//[105 98 99 100 101 102]
	fmt.Println(strToByte)
}
