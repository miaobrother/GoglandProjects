package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["hello"] = "zhangyalei"
	c.TplName = "user/index.html"
}

func (c *UserController) UserInfo() {
	username := c.GetString("username")
	fmt.Println(username)
	age, err := c.GetInt("age")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(age)
	c.Data["hello"] = "this is a userinfo"
	c.TplName = "user/index.html"
}
