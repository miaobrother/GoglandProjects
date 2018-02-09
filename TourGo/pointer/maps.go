package main

import "fmt"

type Sex struct {
	Lat,Long float64
}

var m map[string]Sex

func main() {
	m = make(map[string]Sex)
	m["bell Labs"] = Sex{
		9 ,- 1,
	}
	fmt.Println(m["bell Labs"])
}
