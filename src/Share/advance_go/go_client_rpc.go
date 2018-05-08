package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width,height int
}

func main() {
	//连接远程rpc服务

	rpc,err := rpc.DialHTTP("tcp","127.0.0.1:4321")
	if err != nil{
		log.Fatal(err)
	}
	ret := 0

	err = rpc.Call("Rect.Area",Params{30,40},&ret)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(ret)

	err = rpc.Call("Rect.Area",Params{30,50},&ret)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(ret)

}
