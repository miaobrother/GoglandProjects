package views

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"fmt"
	"awesomeProject/bin"

)




func AsyncMigrate(c *gin.Context)  {
	var (
		result gin.H
	)
	_,connMap,_,err :=ActionMethod()
	if err!=nil{
		result = gin.H{
			"code":1001,
			"msg" : bin.GetMessgae(1001,""),
		}
		c.JSON(http.StatusOK,result)
		return

	}


	source_ip := c.Query("source_ip")
	domName := c.Query("domName")

	dest_ip := c.Query("dest_ip")

	if len(source_ip) >0 && len(domName) >0 && len(dest_ip) >0 {
		fmt.Println("----!!!!",source_ip,domName,dest_ip)

		lock.RLock()
		source_conn,source_exits := connMap[source_ip]
		dest_conn,dest_exits := connMap[dest_ip]
		defer lock.RUnlock()
		if source_exits && dest_exits {
			_,err :=source_conn.LookupDomainByName(domName)


			if err !=nil{
				result = gin.H{
					"code":404,
					"msg" :fmt.Sprintf("主机%s,获取domain：%s失败",source_ip,domName),

				}
				c.JSON(http.StatusOK,result)
				return

			}
			go bin.ServerMigrate(source_ip,dest_ip,domName,source_conn,dest_conn)

			//_,err = dom.Migrate(&dest_conn,libvirt.MIGRATE_LIVE,"","",0)
			//
			//
			//
			//if err !=nil{
			//	result = gin.H{
			//		"code":404,
			//		"msg" :fmt.Sprintf("主机%s,迁移domain：%s失败;失败信息:%s",source_ip,domName,err),
			//
			//	}
			//	c.JSON(http.StatusOK,result)
			//	return
			//
			//}
			//
			//
			//fmt.Println("source_conn",source_conn,dest_conn)
			result = gin.H{
				"code":202,
				"msg" :fmt.Sprintf("主机%s,迁移domain：%s异步处理,",source_ip,domName),

			}
			c.JSON(http.StatusOK,result)
			return


		}else
		{
			result = gin.H{
				"code":404,
				"msg" : "连接不存在,",
			}
			c.JSON(http.StatusOK,result)
			return
		}





	}else {
		result = gin.H{
			"code":404,
			"msg" : "参数缺少！",
		}
		c.JSON(http.StatusOK,result)
		return
	}


}

func AsyncList(c *gin.Context)  {



}