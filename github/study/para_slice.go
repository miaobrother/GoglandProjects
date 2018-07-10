package main

import "fmt"

type VertexOne struct {
	Lat,Long float64
}

var m map[string]VertexOne

func main() {
	m = make(map[string]VertexOne)
	m["Bell Labs"] = VertexOne{
		40.68433,-74.39967,
	}
	fmt.Println(m["Bell Labs"])
}