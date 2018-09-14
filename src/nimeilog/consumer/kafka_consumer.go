package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"strings"
	"time"
)

var(
	wg sync.WaitGroup
)

func main() {
	consumer,err := sarama.NewConsumer(strings.Split("localhost:9092",","),nil)
	if err != nil{
		fmt.Println("got an error :",err)
		return
	}
	partitionList,err := consumer.Partitions("nginx.log")
	if err != nil{
		fmt.Println("got partitionList failed: ",err)
		return
	}
	//fmt.Println("The paratitionList is:",partitionList)

	for partion := range partitionList{
		pc ,err := consumer.ConsumePartition("nginx.log",int32(partion),sarama.OffsetNewest)
		if err != nil{
			fmt.Println("got an error to start consumer for partiuion %d\n",partion,err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages(){
				if len(msg.Value)== 0{
					continue
				}
				fmt.Printf("partition:%d,key:%#v,valuse:%#v\n",msg.Partition,string(msg.Key),string(msg.Value))
			}
		}(pc)
	}
	time.Sleep(time.Second*1000)
	consumer.Close()
}
