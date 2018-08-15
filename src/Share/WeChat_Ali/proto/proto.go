package proto

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	CmdLoginRequest        = 1001
	CmdLoginResponse       = 1002
	CmdRegisterRequest     = 1003
	CmdRegisterResponse    = 1004
	CmdSendMessageRequest  = 1005
	CmdSendMessageResponse = 1006
	CmdBroadMessage        = 1007
)

type ResponseBase struct {
	ErrorNo int    `json:"errorNo"`
	Message string `json:"message"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ResponseBase
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sexname  string `json:"sexname"`
}

type RegisterResponse struct {
	ResponseBase
}

type MessageRequest struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

type MessageResponse struct {
	ResponseBase
}

//前四个字节是  length
//4个字节    cmdno
//body : []byte
func ReadPackage(conn net.Conn) (body []byte, cmd int32, err error) {
	var length int32
	err = binary.Read(conn, binary.BigEndian, &length) //读取长度
	if err != nil {
		if err == io.EOF {
			fmt.Printf(" di yi ge read from conn %v failed error is %v\n", conn, err)
			return
		}
		return

	}

	err = binary.Read(conn, binary.BigEndian, &cmd) //  读取协议号
	if err != nil {
		if err != io.EOF {
			fmt.Printf(" dier ge read from conn %v ,failed error is %v\n", conn, err)
			return

		}
		return

	}
	fmt.Printf("read cmd succ:%d\n", length)
	//var pos int //初始 位置

	var buf []byte = make([]byte, length)

	_, err = io.ReadFull(conn, buf)
	if err != nil {
		if err == io.EOF {
			fmt.Printf("Read body from conn %v, err:%v\n", conn, err)
			return

		}

	}
	body = buf
	return
	/*
		var currPos int32
		for {
			n,errRet := conn.Read(buf)
			if err != nil{
				if errRet != io.EOF{

					return
				}
				body = append(body,buf[:n]...)
				currPos += int32(n)
				if (currPos == length){
					break
				}

				buf = make([]byte,length - currPos)

			}
		}
	*/

	return

}

func WritePackage(conn net.Conn, cmdno int32, body []byte) (err error) {
	var length int32 = int32(len(body))
	err = binary.Write(conn, binary.BigEndian, length)
	if err != nil {
		fmt.Println("Write lenght body failed: ", err)
		return
	}
	fmt.Printf("write length succ:%d\n", length)
	err = binary.Write(conn, binary.BigEndian, cmdno)
	if err != nil {
		fmt.Printf("The write cmdno error is %v", err)
		return
	}
	msglen := len(body) //
	var sendBytes int
	for {

		n, err := conn.Write(body)
		if err != nil {
			fmt.Printf("send to client:%v failed err is %s\n", conn, err)
		}
		sendBytes += n
		if sendBytes >= msglen {
			break
		}
		body = body[sendBytes:]

	}
	return

}
