package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(fileName string, conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err ", err)
	}
	buf := make([]byte, 1024*2)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn err ", err)
			}
			return
		}
		f.Write(buf[:n]) //写入文件
	}
}
func main() {
	//监听
	l, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	defer l.Close()
	//阻塞等待用户连接
	conn, errTwo := l.Accept()
	if errTwo != nil {
		fmt.Println("l.accept err ", errTwo)
		return
	}
	buf := make([]byte, 1024*4)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.read err", err)
		return
	}
	fileName := string(buf[:n])
	fmt.Println(fileName)

	conn.Write([]byte("ok"))

	RecvFile(fileName, conn)
}
