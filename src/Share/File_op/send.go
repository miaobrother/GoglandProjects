package main

import (
	"fmt"
	"os"
	//"io"
	"net"
	//"time"
	"io"
)

func SendFile(path string,conn net.Conn)  {

	// 以只读模式打开文件

	f,err := os.Open(path)
		if err != nil{
			fmt.Println("os.Open", err)
			return
		}
		defer f.Close()
		buf := make([]byte,1024)
		//读取文件读多少发送多少
		for{
			n,err := f.Read(buf)
			if err != nil{
				if err == io.EOF{
					fmt.Println("发送完毕")
				}else {
					fmt.Println("f.read err ",err)
				}
				return
			}
			conn.Write(buf[:n])//给服务器发送
	}

}
func main() {
	// 输入文件
	fmt.Println("Please input file name:")
	var path string
	fmt.Scan(&path)

	info,err := os.Stat(path)
	if err != nil{
		fmt.Println("os.Stat err ",err)
		return
	}

	conn,errOne := net.Dial("tcp","localhost:9000")

	if errOne != nil{
		fmt.Println("存在错误")
		return
	}
	defer conn.Close()


	_,err = conn.Write([]byte(info.Name())) //发送文件名字
	if err != nil{
		fmt.Println("conn.Write err =",err)
		return
	}

	var n int
	buf := make([]byte,1024*1)
	n,err = conn.Read(buf)
	if err != nil{
		fmt.Println("conn read err,",err)
		return
	}
	if "ok"== string(buf[:n]){

		SendFile(path,conn)
	}

}
