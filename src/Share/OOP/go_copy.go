package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3{
		fmt.Println("usage: xxx srcfile dstfile")
		return
	}


	srcfileName := list[1]
	dstfileName := list[2]
	if srcfileName == dstfileName{
		fmt.Println(" 源文件 与目标文件名字异常，请检查")
		return
	}
	//只读方式打开文件
	sf ,err := os.Open(srcfileName)
	if err != nil{
		fmt.Println(err)
		return
	}
	//新建文件
	df,err := os.Create(dstfileName)
	if err != nil{
		fmt.Println(err)
		return
	}
	//close file
	defer sf.Close()
	defer df.Close()


	//逻辑处理
	buf := make([]byte,1024*4)
	for{
		n,err := sf.Read(buf)
		if err != nil{
			if err == io.EOF{
				break
			}
			fmt.Println(err)
		}
		df.Write(buf[:n])
	}
}