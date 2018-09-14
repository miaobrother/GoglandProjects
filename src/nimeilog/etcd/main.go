package main

import (
	"fmt"
	etcd_client "github.com/coreos/etcd/clientv3"
	"time"
)

func main(){
	cli,err := etcd_client.New(etcd_client.Config{//生成一个客户端
		Endpoints:[]string{"localhost:2379"},
		DialTimeout: 1 * time.Second,
	})
	if err != nil{
		fmt.Println("connect failed,error is:",err)
		return
	}
	fmt.Println("connect success")
	defer cli.Close()
}


