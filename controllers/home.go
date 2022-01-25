package controllers

import (
	"net/http"

	"webchat/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	// 检查是否登陆, 若未登陆则跳转到登陆页面
	suser := c.GetSession("user")
	if suser == nil {
		c.Redirect(beego.URLFor("UserController.Login"), http.StatusFound)
	} else {
		c.Data["user"], _ = suser.(models.User)
		c.TplName = "index.html"
	}
}
