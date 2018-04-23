package main

import (
	"fmt"
	"net"
	"log"
	//"reflect"
	"strings"
)

func HandleConn(conn net.Conn)  {
	addr := conn.RemoteAddr()//获取客户端网络地址
	fmt.Println("The addr is successful" ,addr)
	buf := make([]byte,1024*4)

	n,err := conn.Read(buf)
	if err != nil{
		fmt.Println("err = ",err)
		return
	}
	fmt.Printf("[%s] : %s\n",addr,string(buf[:n]))
	fmt.Println("read buf = ",string(buf[:n]))
	if "exit" == string(buf[:n]){
		fmt.Println("主动退出")
		return
	}
	//将数据变化为大写，再给用户发送
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

}

func main() {
	//创建监听
	l ,err := net.Listen("tcp","localhost:9911")
	if err != nil{
		log.Fatal(err)
		return
	}
	defer l.Close()  //最后执行关闭
	//接受多个用户连接
	//阻塞等待用户连接
	//buf := make([]byte,1024*4)
	for{
		conn,err := l.Accept()
		if err != nil{
			fmt.Println("err = ",err)
			return 
		}
		go HandleConn(conn)
		/*
		n,errOne := conn.Read(buf)
		if errOne != nil{
			log.Fatal(errOne)
			continue
		}
		

		fmt.Println("print buf :",string(buf[:n]))
		*/

	}


}

