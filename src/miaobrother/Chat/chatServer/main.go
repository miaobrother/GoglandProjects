package main

import (
	"fmt"
	"net"
	"log"
)

var (
	clientMgr *ClientMgr
)
func startServer(addr string)(l net.Listener,err error)  {
	l,err = net.Listen("tcp",addr)
	if err != nil{
		log.Fatal(err)
		return
	}
	return
}
func main() {
	fmt.Println("Start The Server...")
	clientMgr = NewClientMgr(200)
	l,err := startServer("localhost:9090")
	if err != nil{
		log.Fatal(err)
		return
	}
	err = runServer(l)
	if err !=nil{
		log.Println(err)
		return
	}
	fmt.Println("Server is down")
}
