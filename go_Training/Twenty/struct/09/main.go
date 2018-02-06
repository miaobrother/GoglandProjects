package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	p1 := person{"Zhang","San",20,32}
	json.NewEncoder(os.Stdout).Encode(p1)
	//json.NewDecoder(os.Stdout).Decode(p1)
}
