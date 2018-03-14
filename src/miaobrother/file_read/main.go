package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func testFile()  {
	file ,err := os.Open("/Users/zhangyalei/Desktop/nginx")
	if err != nil{
		fmt.Println("Open file failed:",err)
		return
	}
	defer file.Close()
	var data [1024]byte
	for{
		n,err := file.Read(data[:])
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Printf("read file error:%s\n",err)
		}
		str := string(data[0:n])
		fmt.Println(str)
	}
}

func testBufio()  {
	file,err := os.Open("/Users/zhangyalei/Desktop/nginx")
	if err != nil{
		fmt.Println("Open file failed",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		line,err := reader.ReadString('\n')//按行读
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("Open file failed",err)
			return
		}
		fmt.Println(line)
	}
}
func main() {
	testFile()
	testBufio()

}
