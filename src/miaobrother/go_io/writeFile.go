package main

import (
	"bufio"
	"os"
	"log"
	//"io"
	"fmt"
)

func main() {
	fileTwo,err := os.OpenFile("test.go",os.O_CREATE|os.O_WRONLY,0666)
	if err != nil{
		log.Fatal(err)
	}
	defer fileTwo.Close()
	writer := bufio.NewWriter(fileTwo)

	for i := 0;i < 10;i ++{
		writer.WriteString("Hello World")
	}
	writer.Flush()
	fmt.Println("The file is :",fileTwo)

}
