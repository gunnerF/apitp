package main

import (
	_ "apitp/routers"

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
	beego.SetStaticPath("/down", "download")
	models.Init()
	beego.Run()
}
