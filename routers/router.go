package routers

import (
	"webchat/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login/", &controllers.UserController{}, "*:Login")
	beego.Router("/code/", &controllers.UserController{}, "*:Code")
	beego.Router("/logout/", &controllers.UserController{}, "*:Logout")
	beego.Router("/users/", &controllers.UserController{}, "*:List")
	beego.Router("/ws/msg/", &controllers.WebSocketController{}, "*:Msg")
}
