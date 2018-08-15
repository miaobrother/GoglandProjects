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
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("can't receive")
			break
		}
		fmt.Println("receive back from client:" + reply)

		msg := "receive:" + reply

		fmt.Println("send to client:" + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			break
		}
	}
}

func web(w http.ResponseWriter, r *http.Request) {
	// print method
	fmt.Println("method is", r.Method)
	if r.Method == "GET" {
		tt, _ := template.ParseFiles("websocket.html")
		tt.Execute(w, nil)
	} else {
		fmt.Println(r.PostFormValue("username"))
	}
}

func main() {

	//接受websocket的路由地址
	http.Handle("/websocket", websocket.Handler(Echo))

	http.HandleFunc("/web", web)

	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
