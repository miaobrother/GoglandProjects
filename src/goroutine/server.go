package main

import (
	"fmt"
	"net"
	//"net/http"
	"io"
	"strings"
)

func main() {

	fmt.Println("start server...")
	listen, err := net.Listen("tcp", ":50000") //本地启动 11111端口
	if err != nil {                            //检查是否存在err
		fmt.Println("listen failed,err:", err)
		return
	}
	for { //循环接收客户端请求
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err: ", err)
			continue
		}
		go process(conn) //处理客户端请求
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {

		n, err := conn.Read(buf) //循环读取客户端数据
		if err != nil {
			if err == io.EOF {
				continue

			}
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("The return data is :", strings.ToUpper(string(buf[:n]))) //处理客户端请求
	}
}
