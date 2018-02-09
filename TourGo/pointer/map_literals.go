package main

import "fmt"

type Mac struct {
	Lat,Long float64
}

var m1 = map[string]Mac{
	"Zhang":Mac{
		1,-1,
	},
	"Google":Mac{
		9,-7,
	},
}
func main() {
	fmt.Println(m1)
}