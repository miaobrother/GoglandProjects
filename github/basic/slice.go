package main

import "fmt"


var arr = [10]byte{'a','b','c','d','e','f','k','j','i','g',}
func main() {
	fmt.Println(len(arr))
	slice := arr[4:]
	fmt.Println(slice,len(slice),cap(slice))
	//arr2 := [10]byte{'a'}
	//fmt.Println(arr2)
}
