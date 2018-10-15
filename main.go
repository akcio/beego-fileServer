package main

import (
	_ "fileServer/testBeeGo/routers"
	"github.com/astaxie/beego"
)


func main() {
	beego.BConfig.WebConfig.Session.SessionOn = beego.AppConfig.Bool("sessionOn")
	beego.Run()
}

