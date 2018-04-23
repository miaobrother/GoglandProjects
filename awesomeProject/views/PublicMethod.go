package views

import (
	"awesomeProject/lib"
	"awesomeProject/bin"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"os"
)

func ActionMethod() (actionMaps bin.ControllerMapsType,
	connmaps map[string]libvirt.Connect,MysqlServerMaps map[int]string,
	err error) {
	MysqlServerMap:=make(map[int]string,2096)

	vmservers,err :=lib.SelectVmServer()
	if err!=nil{
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return
	}
	for i:=0;i<len(vmservers);i++{
		ids := vmservers[i].ID
		server_ip := vmservers[i].ServerIp
		MysqlServerMap[ids]  = server_ip
	}
	actionMap,err :=bin.Make()
	if err!=nil{
		fmt.Printf("Make libvirtd conn error:%s",err)
	}
	connmap :=bin.ConnMap
	return actionMap,connmap,MysqlServerMap,err
}
