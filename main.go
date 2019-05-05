package main

import (
	_ "apitp/routers"

	"apitp/commands"
	"apitp/models"
	"apitp/utils"
	"github.com/astaxie/beego"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	//设置静态资源下载目录
	beego.SetStaticPath("/down", "download")
	//初始化model
	models.Init()
	//初始化定时任务
	commands.NewWebSocketTask()
	beego.Run()
}
