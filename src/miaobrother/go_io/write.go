package main

import (
	"os"
	"log"
)

func main() {
	fileOne ,err := os.OpenFile("test.dat",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0666)
	if err != nil{
		log.Fatal(err)
		return
	}

	defer fileOne.Close()
	str := "Hello World"
	fileOne.Write([]byte(str))
	fileOne.WriteString(str)
}