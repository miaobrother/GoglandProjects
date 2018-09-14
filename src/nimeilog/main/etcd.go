package main

import (
	"fmt"
	etcd_client "github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
	"strings"
	"github.com/astaxie/beego/logs"
)

type EtcdClient struct {
	client *etcd_client.Client
}

var(
	etcdClient *EtcdClient
)
func InitEtcd(addr , key string) (err error) {
	cli, err := etcd_client.New(etcd_client.Config{ //生成一个客户端
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 1 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		fmt.Println("connect failed,error is:", err)
		return
	}
	etcdClient = &EtcdClient{
		client:cli,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	if strings.HasSuffix(key,"/") == false{
		key = key + "/"
	}
	etcdKey := fmt.Sprintf("%s%s",key,localIp)
	//_, err = cli.Put(ctx, etcdKey, "simple_value")
	//cancel()
	//if err != nil {
	//	fmt.Println("put ctx simple_value failed.")
	//	return
	//}
	//fmt.Println("Put ctx simple_value success.")

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, etcdKey)
	cancel()
	logs.Debug("resp from etcd:%v\n",resp.Kvs)
	if err != nil {
		fmt.Println("Get failed,err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("key is:%#v valuse is:%#v\n", string(ev.Key), string(ev.Value))
	}
	return

}
