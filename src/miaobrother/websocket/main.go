package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		//websocket接受信息
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("receive back from client:" + reply)
		msg := "receive: " + reply
		fmt.Println("send to client :" + msg)
		//这里是发送消息
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
	return
}

func web(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //打印请求方法
	if r.Method == "Get" {
		t, _ := template.ParseFiles("websocket.html")
		t.Execute(w, nil)
	} else {
		fmt.Println(r.PostFormValue("username"))
		fmt.Println(r.PostFormValue("password"))
	}
}

func main() {
	//接受websocket的路由地址
	http.Handle("/websocket", websocket.Handler(Echo))
	//打开html页面

	http.HandleFunc("web", web)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("LisenServer:", err)
	}
}
