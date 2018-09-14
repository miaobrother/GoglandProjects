package main

import (
	"net"
	"fmt"
)

var(
	localIp string
)
func init()  {

	addrS,_ := net.InterfaceAddrs()
	for _,addr := range addrS{
		if ipnet,ok := addr.(*net.IPNet);ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
				localIp = ipnet.IP.String()
			}
		}
	}
	fmt.Println("The ip is :",localIp)

}

