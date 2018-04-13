package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Server struct {
	ServerName string
	ServerIp string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	//解析一个json

	var s Serverslice

	str := `{"servers":[{"serverName":"GZZ_VPN","serverIP":"127.0.0.1"},{"serverName":"SZ_VPN","serverIp":"127.0.0.2"}]}`
	//json decode
	json.Unmarshal([]byte(str),&s)
	fmt.Println(s)


	s.Servers = append(s.Servers,Server{ServerName:"Shanghai",ServerIp:"127.0.0.3"})
	s.Servers = append(s.Servers,Server{ServerName:"Beijing",ServerIp:"127.0.0.4"})
	b,err := json.Marshal(s)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(b))
}



