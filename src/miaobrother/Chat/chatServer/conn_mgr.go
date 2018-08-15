package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type ClientMgr struct { //定义一个结构体
	clientsMap    map[net.Conn]int //维护所有客户端连接的一个Map
	maxClientNums int
	msgChan       chan []byte //定义一个管道，用来保存客户端发来的消息

	newClientChan chan net.Conn
	closeChan     chan net.Conn
	sy            sync.RWMutex
}

func NewClientMgr(maxClients int) *ClientMgr {
	mgr := &ClientMgr{
		clientsMap:    make(map[net.Conn]int, 1024),
		maxClientNums: maxClients,
		msgChan:       make(chan []byte, 100),
		newClientChan: make(chan net.Conn, 100),
		closeChan:     make(chan net.Conn, 20),
	}
	go mgr.run()
	go mgr.processConn()
	return mgr
}

func (c *ClientMgr) processConn() {
	for {
		select {
		case conn := <-c.newClientChan:
			c.sy.Lock()
			c.clientsMap[conn] = 0
			c.sy.Unlock()
		case conn := <-c.closeChan:
			c.sy.Lock()
			delete(c.clientsMap, conn)
			c.sy.Unlock()

		}

	}
	/*
		for conn := range c.clientsMap{
			c.sy.Lock()
			c.clientsMap[conn] = 0
			c.sy.Unlock()
		}
	*/
}

//遍历所有的客户端发来的消息，并广播到所有的客户端
func (c *ClientMgr) run() {
	for msg := range c.msgChan {
		c.transfer(msg)
	}
}

//广播消息
func (c *ClientMgr) transfer(msg []byte) {
	c.sy.RLock()
	defer c.sy.RUnlock()
	for client, _ := range c.clientsMap {
		err := c.sendToClient(client, msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// 发送消息给指定客户端
func (c *ClientMgr) sendToClient(client net.Conn, msg []byte) (err error) {
	var sendBytes int
	msgLen := len(msg)
	for {
		n, err := client.Write(msg)
		if err != nil {
			log.Fatal(err)
			client.Close()
			delete(c.clientsMap, client)
		}
		sendBytes += n
		if sendBytes >= msgLen {
			break
		}
		msg = msg[sendBytes:]

	}
	return

}

func (c *ClientMgr) addMsg(msg []byte) (err error) {
	ticker := time.NewTicker(time.Millisecond * 10)
	defer ticker.Stop()
	select {
	case c.msgChan <- msg:
		fmt.Printf("send to chan succ\n")
	case <-ticker.C:
		fmt.Printf("time out")
		err = errors.New("The chan s time out")
	}
	return
}
