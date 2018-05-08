package main

import (
	"fmt"
	"net"
	//"go-ethereum/swarm/api/client"
	"errors"
	"sync"
	"time"
)

type ClientMgr struct {
	//clients []net.Conn //所有客户端连接的切片
	clientsMap    map[net.Conn]int //将客户端连接设置为Map 方便查找
	maxclientnums int
	msgChan       chan []byte
	clientChan    chan net.Conn
	closeChan     chan net.Conn
	lock          sync.RWMutex
}

func NewClientMgr(maxclients int) *ClientMgr {
	mgr := &ClientMgr{ //初始化ClientMgr
		clientsMap:    make(map[net.Conn]int, 1024),
		maxclientnums: maxclients,
		msgChan:       make(chan []byte, 1000),
		clientChan:    make(chan net.Conn, 1000),
		closeChan:     make(chan net.Conn, 1000),
	}
	go mgr.run()
	go mgr.ProcConn()
	return mgr
}

func (c *ClientMgr) run() { //遍历所有客户端发送过来的消息，并广播到所有的客户端
	for msg := range c.msgChan {
		c.transfer(msg)
	}
}

func (c *ClientMgr) ProcConn() { //遍历所有客户端发送过来的消息，并广播到所有的客户端
	for {
		select {
		case conn := <-c.clientChan:
			c.lock.Lock()
			c.clientsMap[conn] = 0
			defer c.lock.Unlock()
		case conn := <-c.closeChan:
			c.lock.Lock()
			delete(c.clientsMap, conn)
			defer c.lock.RUnlock()

		}
	}
}

func (c *ClientMgr) sendToClient(client net.Conn, msg []byte) (err error) { //发送消息给指定的客户端
	msglen := len(msg) //
	var sendBytes int
	for {

		n, err := client.Write(msg)
		if err != nil {
			fmt.Printf("send to client:%v failed err is %s\n", client, err)
			client.Close()
			delete(c.clientsMap, client)
		}
		sendBytes += n
		if sendBytes >= msglen {
			break
		}
		msg = msg[sendBytes:]

	}
	return
}
func (c *ClientMgr) transfer(msg []byte) { //发送给当前所有在线用户  又称  广播消息
	c.lock.RLock()
	defer c.lock.RUnlock()
	for client, _ := range c.clientsMap {
		err := c.sendToClient(client, msg)
		if err != nil {
			continue

		}

	}
}

func (c *ClientMgr) addMsg(msg []byte) (err error) { //将msg的消息添加到 chan中
	ticker := time.NewTicker(time.Millisecond * 10)
	defer ticker.Stop()
	select {
	case c.msgChan <- msg:
		fmt.Printf("send to chan succ\n")
	case <-ticker.C:
		fmt.Printf("add msg timeout\n")
		err = errors.New("add msg timeout")

	}
	return
}
