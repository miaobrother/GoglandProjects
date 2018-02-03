package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/zhangyalei/Desktop/test.txt.gz")

	if err != nil {
		fmt.Println("open the failed:", err)
		return
	}

	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("gzip failed,err:%v\n", err)
		return
	}
	bufReader := bufio.NewReader(reader)

	for {
		line, err := bufReader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("read failed,err:%v\n", err)
			return
		}
		fmt.Printf(line)
	}
}
