package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fileServer/testBeeGo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

	o := orm.NewOrm()
	o.Using("default")

	hasher := md5.New()
	hasher.Write([]byte("123456"))
	hashedPass := hex.EncodeToString(hasher.Sum(nil))

	user := models.User{Login: "test", Password: hashedPass}
	id, err := o.Insert(&user)
	if err == nil {
		msg := fmt.Sprintf("Article inserted with id:", id)
		beego.Debug(msg)
	} else {
		msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
		beego.Debug(msg)
	}


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

