package main

import (
	_ "fileServer/testBeeGo/routers"
	"github.com/astaxie/beego"
)


func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

