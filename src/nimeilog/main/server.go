package main

import (
	"github.com/astaxie/beego/logs"
	"nimeilog/kafka"
	"nimeilog/tailf"
	"time"
)

func RunServer() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Warn("send to kafka failed,error is: %v\n", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return

}
