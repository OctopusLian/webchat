package controllers

import (
	"fmt"
	"net/http"
	"sort"

	"webchat/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	if c.Ctx.Input.IsGet() {
		c.TplName = "login.html"
	} else if c.Ctx.Input.IsPost() {
		// 定义response信息
		result := struct {
			Code int `json:"code"`
		}{400}

		// 获取提交数据
		email := c.GetString("email")
		code := c.GetString("code")

		// 登陆验证
		if user, err := models.Authenticate(email, code); err == nil {
			// 验证成功
			result.Code = 200
			c.SetSession("user", *user)
		}
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		c.CustomAbort(http.StatusMethodNotAllowed, "")
	}
}

func (c *UserController) Code() {
	// 定义response信息
	result := struct {
		Code int `json:"code"`
	}{400}

	email := c.GetString("email")

	if err := models.CreateLoginCode(email); err == nil {
		result.Code = 200
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *UserController) Logout() {
	c.DestroySession()
	c.Redirect(beego.URLFor("UserController.Login"), http.StatusFound)
}

func (c *UserController) List() {
	// 定义response信息
	result := struct {
		Code  int             `json:"code"`
		Users []models.Status `json:"users"`
	}{400, nil}
	suser := c.GetSession("user")
	fmt.Println(suser)
	if suser != nil {
		result.Code = 200
		for _, user := range models.AllUser() {
			_, exists := clients[user.Id]
			result.Users = append(result.Users, models.Status{exists, user})
		}
		sort.Slice(result.Users, func(i, j int) bool {
			if result.Users[i].Online == result.Users[j].Online {
				return result.Users[i].User.Id > result.Users[j].User.Id
			} else if result.Users[i].Online {
				return true
			} else {
				return false
			}
		})
	}
	c.Data["json"] = result
	c.ServeJSON()
}
