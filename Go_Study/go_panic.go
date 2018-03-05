package main

import (
	"os"
	//"fmt"
)

func main() {
	panic("a problem.")

	_,err := os.Create("lastname")
	if err != nil{
		panic(err)
	}
}