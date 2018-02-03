package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 13
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	delete(m, "key1")
	fmt.Println(m)

	k, prs := m["key2"]
	fmt.Println("k", k, "prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
