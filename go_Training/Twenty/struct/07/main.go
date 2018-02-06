package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	p1 := person{"Zhang","Yalei",19,111}

	bs ,_ := json.Marshal(p1)
	fmt.Println(bs)

	fmt.Printf("%T\n",bs)

	fmt.Println(string(bs))
}