package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		fmt.Println("producer close,error is : ", err)
		return
	}
	return

}

func SendToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid, offset, errs := client.SendMessage(msg)
	if errs != nil {
		logs.Error("send message failed,err:%v,data:%v,topic is:%v\n", err, data, topic)
		return
	}
	time.Sleep(time.Second*1)
	logs.Debug("send success pid:%d,offset is:%v,topic is %v\n", pid, offset, topic)
    return
}
