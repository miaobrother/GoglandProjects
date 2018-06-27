package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a :=names[0:2]
	var b = names[2:3]
	fmt.Println(b)
	b[0] = "nimeia"
	fmt.Println(a)
	fmt.Println(names)
	n := []struct {
		i int
		b bool
	}{
		{2,true},
		{3,false},
	}
	fmt.Println(n[0].i)

}
