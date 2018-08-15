package routers

import (
	"garage_new/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/Postuser", &controllers.UserController{}, "*:UserInfo")
}
