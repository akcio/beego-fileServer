package controllers

import (
	"github.com/astaxie/beego"
)

type NewController struct {
	beego.Controller
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = main.Ctx.Input.Param(":id")

	main.TplName = "default/hello-sitepoint.tpl"
}

func (this *NewController) Get() {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)

	var val = sess.Get("test")
	if val != "10"{
		this.Data["Val"] = 10
		sess.Set("test", "10")
	} else {
		sess.Set("test", "vall")
		this.Data["Val"] = 50

	}
	this.Data["Files"] = "file1"
	this.TplName = "fileList.tpl"

}