package main

import (
	//"fmt"
	"net"
	"log"
)

func main() {
	conn,err := net.Dial("tcp","localhost:9911")
	if err != nil{
		log.Fatal(err)
		return
	}
	defer conn.Close()

	//发送数据
	conn.Write([]byte("what's going on"))
}
