package main

import "fmt"

type TZZ int

func (tz *TZZ) Increase(num int) {
	*tz += TZZ(num)
}

func main() {
	var a TZZ
	a.Increase(100)
	fmt.Println(a)
}
