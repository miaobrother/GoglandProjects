package main

import (
	"fmt"
	//"io"
	"Share/WeChat_Ali/proto"
	"encoding/json"
	"net"
	//"github.com/gin-gonic/gin/json"
	"io"
)

func Process(conn net.Conn) { // 处理客户端发来的消息

	defer func() {
		clientMgr.closeChan <- conn
		conn.Close()
	}()
	for {
		body, cmd, err := proto.ReadPackage(conn)
		if err != nil {
			if err != io.EOF{
				fmt.Printf("The conn is %v error is %v\n", conn, err)
				return

			}
			return

		}
		err = processRequest(conn, body, cmd)
		if err != nil {
			fmt.Printf("processRequest[%v] failed,err:%v\n", cmd, err)
			return
		}

		/*
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Printf("read from client %#v failed ,The err is %v\n", conn, err)
				}

				return
			}
			buf = buf[:n]
			clientMgr.addMsg(buf)
		*/
	}
}

func RunServer(listener net.Listener) (err error) {
	fmt.Println("Run successful server")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("The Recv failed:", err)
			continue
		}
		clientMgr.clientChan <- conn
		go Process(conn)
	}
}

func processRequest(conn net.Conn, body []byte, cmd int32) (err error) {
	switch cmd {
	case proto.CmdLoginRequest:
		err = processLogin(conn, body)
	case proto.CmdRegisterRequest:
		err = processRegister(conn, body)
	case proto.CmdSendMessageRequest:
		err = processMessage(conn, body)
	default:
		fmt.Printf("Unsupport [%v]request", cmd)
		return
	}
	return
}

func processLogin(conn net.Conn, body []byte) (err error) {
	var loginRequest proto.LoginRequest
	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		fmt.Printf("The body unmarshal failed error is %v\n", err)
		return
	}
	var loginResp proto.LoginResponse
	if loginRequest.Password == "zhangsan" && loginRequest.Password == "miaozi" {
		loginResp.ErrorNo = 0
		loginResp.Message = "Successful"
		return
	}
	loginResp.ErrorNo = 100
	loginResp.Message = "Username or Passwd not vailed"
	data, err := json.Marshal(loginResp)
	if err != nil {
		fmt.Printf("Marshal failed[%v]\n", err)
		return
	}
	return proto.WritePackage(conn, proto.CmdLoginResponse, data)
}

func processRegister(conn net.Conn, body []byte) (err error) {
	return
}

func processMessage(conn net.Conn, body []byte) (err error) {
	return
}
