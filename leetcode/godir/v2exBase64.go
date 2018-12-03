package main

import (
	"encoding/base64"
	"log"
	"fmt"
)

const encodeing  = "NjY1MTEyMjU4"
func main() {
	decoded,err := base64.StdEncoding.DecodeString(encodeing)
	if err != nil{
		log.Fatal("Got a error: ",err)
		return
	}
	fmt.Println(string(decoded))

}
