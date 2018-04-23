package main

import (
	"fmt"
	"net"
	//"go-ethereum/les"
)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

var onlineMap map[string]Client

var message = make(chan string)

func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息前，这里会阻塞
		//range map 给每个map成员发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg

		}

	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": login succ" +":" +msg
	return
}
func HandleConn(conn net.Conn) { //处理用户连接
	defer conn.Close()
	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()
	//创建一个结构体，默认 用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = cli
	// 新开一个协程，专门处理给当前客户端发送信息

	go WriteMsgToClient(cli, conn)
	//广播某个在线
	//message <- "[" +cli.Addr +"]"+ cli.Name +": login succ"
	message <- MakeMsg(cli,"login")
	//创建一个匿名函数，接收用户发送过来的数据
	go func() {
		buf := make([]byte, 1024*2)
		for {

			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Println("conn.Read err", err)
				return
			}
			msg := string(buf[:n-1])
			if len(msg) == 3 && msg == "who"{
				//遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _,tmp  := range onlineMap{
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}

			}else {
				//转发内容
				message <- MakeMsg(cli,msg)
			}
			//message <- MakeMsg(cli, msg)
		}
	}()
	for{

	}
}

func main() {
	//监听
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	defer l.Close()

	go Manager()
	//新开一个协程，转发消息，只要消息来了，会通过map给所有成员发消息
	//主协程 阻塞等待用户连接
	for {
		conn, err := l.Accept() //不停的接收用户发来的请求
		if err != nil {
			fmt.Println("conn err is ", err)
			continue
		}
		go HandleConn(conn) // 处理用户连接

	}

}
