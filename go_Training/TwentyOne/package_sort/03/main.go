package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"Zhang","San","li","si"}
	sort.Strings(s)
	fmt.Println(s)
}