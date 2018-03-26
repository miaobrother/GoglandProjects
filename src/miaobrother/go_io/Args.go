package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println("args[0]= ",os.Args[0])
	if len(os.Args) > 1{
		for index,v := range os.Args{
			if index == 0{
				continue
			}
			fmt.Println(v)
		}
	}
}