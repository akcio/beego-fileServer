package main

import (
	_ "beego-fileServer/routers"
	"github.com/astaxie/beego"
)


func main() {
	beego.BConfig.MaxMemory = 1 << 22
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

