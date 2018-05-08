package main

import (
	"fmt"
	"log"
	"net"
)

var clientMgr *ClientMgr

func StartServer(addr string) (listener net.Listener, err error) {
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("listen addr is %s and failed is %v\n", addr, err)
		return
	}
	return
}

func main() {
	clientMgr = NewClientMgr(1000)
	fmt.Println("Start server ...")
	listener, err := StartServer(":8080")
	if err != nil {
		log.Fatal("The error is", err)
		return
	}

	err = RunServer(listener)
	if err != nil {
		fmt.Println("run server failed,err:", err)
		return
	}
}
