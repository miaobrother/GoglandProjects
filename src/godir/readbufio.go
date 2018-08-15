package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	//"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes,err := reader.ReadString('\n')
	if err != nil{
		log.Fatal("the error is",err)
		return
	}
	fmt.Println(bytes)
}
