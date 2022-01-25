/*
 * @Author: zhangniannian
 * @Date: 2022-01-25 14:49:45
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-25 15:11:30
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	_ "webchat/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local",
		beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpassword"),
		beego.AppConfig.String("dbhost"),
		beego.AppConfig.String("dbport"),
		beego.AppConfig.String("dbname"),
	)
	orm.RegisterDataBase("default", "mysql", driverSource) //注册数据库

	orm.RunSyncdb("default", false, true)
	// 启动beego 服务
	beego.Run()
}
