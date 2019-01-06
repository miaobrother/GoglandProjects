package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("the error is", err)
		return
	}
	fmt.Println(bytes)
}
