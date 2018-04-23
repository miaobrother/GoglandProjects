package main


import (
	"fmt"
	"awesomeProject/bin"

	"github.com/gin-gonic/gin"
	"awesomeProject/views"
)
var (
	conn_obj *bin.ConnMain

	)

func init()  {
	conn_obj =bin.NewMgr()
	bin.InitAsync()
	go conn_obj.Run()
	fmt.Println("开始获取libvirtd对象",)
}



func main()  {
	defer func() {
		err :=recover()
		if err !=nil{
			fmt.Println(err)
		}
	}()


	router := gin.Default()
	group1 :=router.Group("/api/v1/")
	group2 :=router.Group("/api/v1/")
	group3 :=router.Group("/api/v2/")
	// 最基本的用法
	//注意':'必须要匹配,'*'选择匹配,即存在就匹配,否则可以不考虑

	{
		//http://192.168.3.79:8888/api/v1/libvirt/action?server_id=30&vm_name=别克微信商城测试-10.1.11.236&op=Shutdown
		group1.POST("/:libvirt/action", views.Action)
		//http://10.1.11.24:8888/api/v1/libvirt/getdominfo?server_id=6&action=GetDomainINfo
	}

	{
		group2.GET("/:libvirt/getdominfo", views.GetDomainInfo)
	}
	{
		group3.POST("/:migrate/async", views.AsyncMigrate)
		group3.GET("/:migrate/asynclist", views.AsyncList)

	}

	// 绑定端口是8888
	router.Run(":8888")

}