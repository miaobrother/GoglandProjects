package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	data := make([]byte,10000)
	file ,err := os.Open("channel.go")
	if err != nil{
		log.Fatal(err)
	}
	count,err := file.Read(data)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes:%q\n",count,data[:count])
}