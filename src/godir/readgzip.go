package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("abc.tar.gz")
	if err != nil {
		log.Fatal("Got a error: ", err)
		return
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal("Got a Gzip error: ", err)
		return
	}
	//TODO : 去除 '\n'
	buferRead := bufio.NewReader(reader)
	for {
		line, err := buferRead.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Got a error is:", err)
			return
		}
		fmt.Printf("The line is :%v\n", strings.Replace(line,"\n","",-1))

	}

}
