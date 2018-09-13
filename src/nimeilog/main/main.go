package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"nimeilog/kafka"
	"nimeilog/tailf"
)

func main() {
	filename := "/Users/zhangyalei/GoglandProjects/src/nimeilog/conf/logconf.conf"
	err := LoadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed,err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = InitLogger()
	if err != nil {
		fmt.Printf("load logger failed,err:%v\n", err)
		panic("load logger failed")
		return
	}

	err = tailf.InitTail(appConfig.collConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tailf failed,err is :%v\n", err)
		return
	}

	err = kafka.InitKafka(appConfig.kafka)
	if err != nil {
		logs.Error("initkafka failed error is:%v\n", err)
		return
	}

	logs.Debug("initialize all success")
	logs.Debug("got conf is: %v\n", appConfig)
	/*
		go func() {
			for {
				rand.Seed(time.Now().Unix())
				time.Sleep(time.Second * 2)
				fmt.Println("This is a test log ", rand.Intn(10))
			}
		}()
	*/

	err = RunServer()
	if err != nil {
		logs.Error("server run failed,err is:", err)
		return
	}
	logs.Info("program exit")
}
