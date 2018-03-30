package proto

import (
	"net"
	"encoding/binary"
	"log"
	"fmt"
	//"github.com/astaxie/beego"
)
const (
	CmdLoginRequest = 1001
	CmdLoginResponse = 1002
	CmdRegisterRequest = 1003
	CmdRegisterResponse = 1004
	CmdSendMessageRequest = 1005
	CmdSendMessageResponse = 1006
	CmdBroadMessage = 1007
)
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ReponseBasic
}


type RegisterReponse struct {
	ReponseBasic
}

type ReponseBasic struct {
	Error int `json:"error"`
	Message string `json:"message"`
}
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sex string `json:"sex"`
}

type MessageRequest struct {
	Message string `json:"message"`
	Username string `json:"username"`
}

type MessageResponse struct {
	ReponseBasic
}

func ReadPacket(conn net.Conn)(body []byte,cmd int32,err error)  {
	var length int32
	err = binary.Read(conn,binary.BigEndian,&length)
	if err != nil{
		fmt.Printf("read from conn:%v failed,err:%v\n",conn,err)
		return
	}

	err = binary.Read(conn,binary.BigEndian,&cmd)
	if err != nil{
		fmt.Printf("read from conn:%v failed,err:%v\n",conn,err)
		return
	}
	var buf []byte =make([]byte,length)
	var curReadBytes int32
	for{
		n,errRet := conn.Read(buf)
		if errRet != nil{
			err = errRet
			log.Fatal(err)
			return
		}
		body = append(body,buf[0:n]...)
		curReadBytes +=  int32(n)
		if curReadBytes == length{
			break
		}

		buf = make([]byte,length - curReadBytes)
	}

	return

}


func WritePacket(conn net.Conn,cmdno int32, body []byte)(err error)  {
	var length int32 = int32(len(body))
	err = binary.Write(conn,binary.BigEndian,length)
	if err != nil{
		log.Fatal(err)
		return
	}
	err = binary.Write(conn,binary.BigEndian,cmdno)

	if err != nil{
		log.Fatal(err)
		return
	}
	var n int
	var sendBytes int
	msgLen := len(body)
	for{
		n,err = conn.Write(body)
		if err  != nil{
			log.Fatal(err)
			return
		}
		sendBytes +=n
		if sendBytes >= msgLen{
			break
		}
		body = body[sendBytes:]
	}
	return

}