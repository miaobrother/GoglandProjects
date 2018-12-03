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
	file, err := os.Open("/Users/zhangyalei/Desktop/error.log-20180824.gz")
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
		newLine := strings.Replace(line,"\n","",-1)
		fmt.Println(strings.Contains(newLine,"http://test.usnoon.com/times/admin/traffic/unregistercarstatis.html?menuid=116"))
		//s := strings.Split(newLine," ")
		//fmt.Println(s)

	}


}
