package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testHandlefunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"login failed": "success",
	})

}
func userInfo(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s", name)

}
func actionInfo(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func welcomInfo(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
func main() {

	r := gin.Default()
	//gin.DisableConsoleColor()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", testHandlefunc)
	r.GET("/user/:name", userInfo)
	r.GET("/user/:name/*action", actionInfo)
	r.GET("/welcome", welcomInfo)
	r.Run()
}
