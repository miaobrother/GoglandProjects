package main

import (
	"fmt"
	"compress/gzip"
	"os"
	//"bufio"
	"io"
	"log"
)

func main() {
	file,err := os.Open( "/Users/zhangyalei/123.tar.gz")
	if err != nil{
		log.Fatal(err,'\n')
		return
	}
	defer file.Close()
	reader,err := gzip.NewReader(file)
	if err != nil{
		fmt.Println("gzip file is failed....")
		return
	}
	var buff[128]byte
	var write[]byte
	for{
		read,err := reader.Read(buff[:])
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("read file:",err)
			return
		}
		write = append(write, buff[:read]...)
	}
	fmt.Println(string(write))
	
}
