package main

import (
	_ "beego-fileServer/routers"
	"github.com/astaxie/beego"
)


func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

