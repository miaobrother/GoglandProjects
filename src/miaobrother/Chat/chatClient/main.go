package main

import (
	"fmt"
	"miaobrother/Chat/proto"
	"net"
	"log"
	"encoding/json"
)

var recvMsg chan interface{}
func main() {
	conn,err := net.Dial("tcp","192.168.100.201")

	if err != nil{
		fmt.Printf("dial server failed,err:%v\n",err)
		return
	}
	recvMsg  = make(chan interface{},1000)
	defer conn.Close()
	go read(conn)
	err = login(conn)
	if err != nil{
		fmt.Printf("login failed,err:%v\n",err)
		return
	}
	msg := <- recvMsg
	loginResp,ok := msg.(*proto.LoginResponse)
	if !ok{
		fmt.Printf("unexpect msg:%T,%+v\n",msg,msg)
		return
	}
	if loginResp.Error != 0{
		fmt.Printf("login failed err is:%v\n",loginResp.Message)
		return
	}
	fmt.Println("登陆成功...")
	for{
		break
	}
}
func login(conn net.Conn)(err error)  {
	var loginReq proto.LoginRequest
	loginReq.Password = "admin"
	loginReq.Username = "admin"

	body,err := json.Marshal(loginReq)
	if err != nil{
		log.Fatal(err)
	}
	err = proto.WritePacket(conn,proto.CmdLoginRequest,body)
	if err != nil{
		fmt.Printf("Send to server failed,err:%v\n",err)
		return
	}
	return

}
func read(conn net.Conn)  {
	for{
		body,cmd,err := proto.ReadPacket(conn)
		if err != nil{
			log.Fatal(err)
			return
		}
		switch cmd {
		case proto.CmdLoginResponse:
			err = processLoginResponse(conn,body)
		case proto.CmdSendMessageResponse:
			err = processSendMsgReponse(conn,body)
		case proto.CmdRegisterResponse:
			err = processBroadMessage(conn,body)
		default:
			fmt.Printf("unspport cmd[%v]\n",cmd)
			//fmt.Println("Login succ ")
			return
		}

	}
}

func processLoginResponse(conn net.Conn,body []byte)(err error)  {
	var loginResponse proto.LoginResponse
	err =  login(conn)
	if err != nil{
		fmt.Printf("login failed %v\n",err)
		return
	}
	recvMsg <- &loginResponse
	return
}

func processSendMsgReponse(conn net.Conn,body []byte)(err error)  {
	return
}


func processBroadMessage(conn net.Conn,body []byte)(err error)  {
	return
}