package views

import (

	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"awesomeProject/bin"
	"strconv"
	"sync"
	"reflect"

)

var (

	lock sync.RWMutex
)



func Action(c *gin.Context)  {
	var (
		result gin.H
	)
	actionMap,connMap,MysqlServerMap,err :=ActionMethod()
	if err!=nil{
		result = gin.H{
			"code":1001,
			"msg" : bin.GetMessgae(1001,""),
		}
		c.JSON(http.StatusOK,result)
		return

	}
	server_id := c.Query("server_id")
	domian_name := c.Query("vm_name")
	action := c.Query("op")

	server_ids,err :=strconv.Atoi(server_id)
	if err!=nil{
		result = gin.H{
			"code":1002,
			"msg" : bin.GetMessgae(1002,server_id),
		}
		c.JSON(http.StatusOK,result)
		return
	}

	ips,exits := MysqlServerMap[server_ids]
	if exits==false{
		result = gin.H{
			"code":1003,
			"msg" : bin.GetMessgae(1003,server_ids),
		}
		c.JSON(http.StatusOK,result)
		return
	}
	lock.RLock()
	conn,exits := connMap[ips]
	lock.RUnlock()
	defer func() {
		err :=recover()
		if err !=nil{
			fmt.Printf("Action views error:%s\n",err)

		}
	}()
	if exits == false {
		result = gin.H{
			"code":1004,
			"msg" : bin.GetMessgae(1004,ips),
		}
		c.JSON(http.StatusOK,result)
		return
	}


	fmt.Println("Action log",server_id,domian_name,action)

	_,exits = actionMap[action]
	if exits == false{
		result = gin.H{
			"code":1005,
			"msg" : bin.GetMessgae(1005,action),
		}
		c.JSON(http.StatusOK,result)
		fmt.Println("test",result)
		return
	}
	fmt.Println("Action log actionMap" ,actionMap)
	parms := []reflect.Value{reflect.ValueOf(&conn),reflect.ValueOf(domian_name)}
	ch := make(chan interface{},2)
	go func() {
		callback:=actionMap[action].Call(parms)
		ch <- callback[0].FieldByName("Code").Int()
		ch <- callback[0].FieldByName("Message").String()
	}()
	Code :=<-ch
	Messgae :=<-ch
	result = gin.H{
		"code":&Code,
		"msg" : &Messgae,
	}
	c.JSON(http.StatusOK,result)

}

func GetDomainInfo(c *gin.Context){
	var (
		result gin.H
	)
	actionMap,connMap,MysqlServerMap,err :=ActionMethod()
	if err!=nil{
		result = gin.H{
			"code":1001,
			"msg" : bin.GetMessgae(1001,""),
		}
		c.JSON(http.StatusOK,result)
		return

	}
	server_id := c.Query("server_id")
	action := c.Query("action")
	server_ids,err :=strconv.Atoi(server_id)
	if err!=nil{
		result = gin.H{
			"code":1002,
			"msg" : bin.GetMessgae(1002,server_id),
		}
		c.JSON(http.StatusOK,result)
		return
	}

	ips,exits := MysqlServerMap[server_ids]
	if exits==false{
		result = gin.H{
			"code":1003,
			"msg" : bin.GetMessgae(1003,server_ids),
		}
		c.JSON(http.StatusOK,result)
		return
	}
	lock.RLock()
	conn,exits := connMap[ips]
	lock.RUnlock()
	defer func() {
		err :=recover()
		if err !=nil{
			fmt.Printf("GetDomainInfo views error:%s\n",err)

		}
	}()

	if exits == false {
		result = gin.H{
			"code":1004,
			"msg" : bin.GetMessgae(1004,ips),
		}
		c.JSON(http.StatusOK,result)
		return
	}

	_,exits = actionMap[action]
	if exits == false{
		result = gin.H{
			"code":1005,
			"msg" : bin.GetMessgae(1005,""),
		}
		c.JSON(http.StatusOK,result)
		return
	}

	parms := []reflect.Value{reflect.ValueOf(&conn),reflect.ValueOf(&ips)}

	ch := make(chan interface{},2)
	go func() {
		callback:=actionMap[action].Call(parms)

		ch <- callback[0].FieldByName("Code").Int()
		ch <- callback[0].FieldByName("Message").Interface()
	}()
	Code :=<-ch
	Messgae :=<-ch
	result = gin.H{
		"code":&Code,
		"msg" : &Messgae,
	}
	c.JSON(http.StatusOK,result)
}
