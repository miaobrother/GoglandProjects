package main

import (
	"fmt"
	"Share/WeChat_Ali/proto"
	"net"
	"encoding/json"
	//"github.com/gin-gonic/gin/json"
)
var recvMsg chan interface{}

func main() {
	recvMsg = make(chan interface{},1000)
	conn,err := net.Dial("tcp",":8080")
	if err != nil{
		fmt.Printf("dial server failed, err is %v\n",err)
		return
	}

	defer conn.Close()
	go read(conn)
	err = login(conn)
	if err != nil{
		fmt.Printf("login failed err is %v ",err )
		return
	}
	msg := <- recvMsg
	loginResp,ok := msg.(*proto.LoginResponse)
	if !ok{
		fmt.Printf("unexpect msg:%T,%+v\n",msg,msg)
		return
	}
	if loginResp.ErrorNo != 0{
		fmt.Printf("login failed,err:%v\n",loginResp.Message)
		return
	}
	fmt.Printf("login success\n")

	for{
		break
	}
}

func login(conn net.Conn)(err error)  {
	var loginReq proto.LoginRequest
	loginReq.Password = "zhangsan"
	loginReq.Username = "miaozi"
	body,err := json.Marshal(loginReq)
	if err != nil{
		fmt.Printf("marshal failed, err: %v\n",err)
		return
	}
	if err != nil{
		fmt.Printf("Marshal failed %v\n",err)
		return
	}
	err = proto.WritePackage(conn,proto.CmdLoginRequest,body)
	if err != nil{
		fmt.Printf("send to server failed errors is : %v\n",err)
		return
	}
	return

}
func read(conn net.Conn)  {
	for{
		body,cmd,err := proto.ReadPackage(conn)
		if err != nil{
			fmt.Printf("read from server failed error is %v\n",err)
			return
		}
		switch cmd {
		case proto.CmdLoginResponse:
			err = processLoginResponse(conn,body)
		case proto.CmdSendMessageResponse:
			err = processSendMsgResponse(conn,body)
		case proto.CmdBroadMessage:
			err = processBroadMessage(conn,body)
		default:
			fmt.Printf("unsupport cmd %v\n",err)
			return
		}
	 }
}

func processLoginResponse(conn net.Conn,body []byte)(err error)  {
	var loginResponse proto.LoginResponse
	err = json.Unmarshal(body,&loginResponse)
	if err != nil{
		fmt.Printf("unmarshal failed, error is %v\n",err)
		return
	}
	recvMsg <- &loginResponse
	return
}


func processSendMsgResponse(conn net.Conn,body []byte)(err error)  {
	return
}

func processBroadMessage(conn net.Conn,body []byte)(err error)  {
	return
}
