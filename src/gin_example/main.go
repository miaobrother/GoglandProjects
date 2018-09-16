package main

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/logs"
	"net/http"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(200,gin.H{
			"message" :"pong",
			"username" :username,
			"address" : address,
		})
	})
	r.POST("/upload", func(c *gin.Context) {
		file,err := c.FormFile("file")
		if err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"message":err.Error(),
			})
			return
		}
		logs.Debug("upload %v failed.error is:%v\n",file,err)
		dst := fmt.Sprintf("/tmp/%v",file.Filename)
		c.SaveUploadedFile(file,dst)
		c.JSON(http.StatusOK,gin.H{
			"message":fmt.Sprintf("file %v upload succes\n",file.Filename),
		})
	})
	r.Run(":9090")
}