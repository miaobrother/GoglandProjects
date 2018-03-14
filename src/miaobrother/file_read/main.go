package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"io/ioutil"
	"compress/gzip"
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

func testIoutil()  {
	data,err := ioutil.ReadFile("/Users/zhangyalei/Desktop/nginx")
	if err != nil{
		fmt.Printf("read file is failed %s",err)
		return
	}
	fmt.Println(string(data))
}
func readGzip()  {
	file,err := os.Open("/Users/zhangyalei/Desktop/nginx.tar.gz")
	if err != nil{
		fmt.Printf("The file is open failed,%s",err)
		return
	}
	defer file.Close()

	reader ,err := gzip.NewReader(file)
	if err != nil{
		fmt.Printf("gzip failed,err:%s",err)
		return
	}
	bufreader := bufio.NewReader(reader)
	for{
		line,err := bufreader.ReadString('\n')
		if err != nil{
			if err == io.EOF{
				break
			}
			fmt.Printf("The file is open failed,%s",err)
			return
		}
		fmt.Printf("%s",line)
	}
}
func main() {
	//testFile()
	//testBufio()
	testIoutil()

}
