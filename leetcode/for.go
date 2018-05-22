package main

import "fmt"

func length(s string) int  {
	fmt.Println("call length.")
	return len(s)
}

func main() {
	var s string = "abc"
	for i,n := 0,length(s); i < n; i++{
		fmt.Println(i,s[i])
	}
}