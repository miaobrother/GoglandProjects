package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"bufio"
)

func testFileRead() {
	file, err := os.Open("/Users/zhangyalei/Desktop/test.txt")

	if err != nil {
		fmt.Println("Open the file %v is:", err)
		return
	}
	var data [1024]byte
	for {
		n, err := file.Read(data[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read the file error:", err)

			return
		}
		str := string(data[0:n])
		fmt.Println(str)
	}
}

func testBufIo() {
	file, err := os.Open("/Users/zhangyalei/Desktop/test.txt")

	if err != nil {
		fmt.Println("open the failed:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file error:%s\n", err)
			return
		}
		fmt.Println(line)
	}
}

func testIoUtil() {
	data, err := ioutil.ReadFile("/Users/zhangyalei/Desktop/test.txt")
	if err != nil {
		fmt.Println("read file failed:%v\n", err)
		return
	}
	fmt.Println("%s\n", string(data))
}

func main() {
	//testFileRead()
	//testBufIo()
	testIoUtil()
}
