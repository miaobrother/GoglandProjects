package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `from:"user" json:"user" binding:"required"`
	Password string `from:"password" json:"password" binding:"required"`
}

func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}

func submit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submit success",
	})
}

func info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "info success",
	})
}
func main() {
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		logs.Debug("upload %v failed.error is:%v\n", file, err)
		dst := fmt.Sprintf("/tmp/%v", file.Filename)
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("file %v upload succes\n", file.Filename),
		})
	})

	r.POST("/multiupload", func(c *gin.Context) {
		from, _ := c.MultipartForm()
		files := from.File["file"]
		for _, value := range files {
			fmt.Println(value.Filename)
			dst := fmt.Sprintf("/tmp/%s_%d", value.Filename, value.Size)
			c.SaveUploadedFile(value, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d file upload", len(files)),
		})

	})

	v1 := r.Group("/v1")
	{
		v1.POST("/logout", logout)
		v1.POST("/submit", submit)
		v1.POST("/user/info", info)
	}

	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		type Msgmsg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		var Msg Msgmsg
		Msg.Name = "zhangsan"
		Msg.Message = "hello zhangsana"
		Msg.Number = 100
		c.JSON(http.StatusOK, Msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		type xmlMsg struct {
			Name  string
			Total int
			Bal   int
			Type  int
		}
		var msg xmlMsg
		msg.Name = "lisi"
		msg.Type = 1
		msg.Total = 100
		msg.Bal = 1000
		c.XML(http.StatusOK, msg)
	})

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "Posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"users/index.html",gin.H{
			"title":"This is a user index",
		})
	})
	r.Run(":9090")
}
