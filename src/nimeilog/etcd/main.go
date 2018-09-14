package main

import (
	"fmt"
	etcd_client "github.com/coreos/etcd/clientv3"
	"time"
	"golang.org/x/net/context"
)

func main(){
	cli,err := etcd_client.New(etcd_client.Config{//生成一个客户端
		Endpoints:[]string{"localhost:2379"},
		DialTimeout: 1 * time.Second,
	})
	defer cli.Close()
	if err != nil{
		fmt.Println("connect failed,error is:",err)
		return
	}
	//fmt.Println("connect success")

	ctx,cancel := context.WithTimeout(context.Background(),1*time.Second)
	_,err = cli.Put(ctx,"../conf/logconf.conf","simple_value")
	cancel()
	if err != nil{
		fmt.Println("put ctx simple_value failed.")
		return
	}
	fmt.Println("Put ctx simple_value success.")

	ctx,cancel = context.WithTimeout(context.Background(),time.Second)
	resp,err := cli.Get(ctx,"../conf/logconf.conf")
	cancel()
	if err != nil{
		fmt.Println("Get failed,err:",err)
		return
	}
	for _, ev := range resp.Kvs{
		fmt.Printf("key is:%v valuse is:%v\n",string(ev.Key),string(ev.Value))
	}

}


