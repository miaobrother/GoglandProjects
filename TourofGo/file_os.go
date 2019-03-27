package main

import (
	"os"
	"log"
	"fmt"
)


func main() {
	data := make([]byte,100)
	file ,err := os.Open("defer.go")
	if err != nil{
		log.Fatal(err)
	}
	count,err := file.Read(data)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("count",count,data[:count])

}