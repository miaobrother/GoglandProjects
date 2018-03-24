package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"io"
)

func main() {
	//var str string
	inPut,err := os.Open("./input.go")
	if err != nil{
		log.Fatal(err)
		return
	}
	defer inPut.Close()
	reader := bufio.NewReader(inPut)
	for{
		strone,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}else if err != nil{
			log.Fatal(err)
			return
		}
		fmt.Println(strone)
	}



}