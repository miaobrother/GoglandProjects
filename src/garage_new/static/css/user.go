package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["hello"] = "zhangyalei"
	c.TplName = "user/index.html"
}