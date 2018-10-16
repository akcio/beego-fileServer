package helpers

import (
	"beego-fileServer/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func SetLayoutFor(this * beego.Controller) {
	this.Layout = "layouts/main_layout.tpl"
	this.Data["Tittle"] = "File server"
	this.Data["IsLogined"] = IsUserLogedIn(this)
}

func GetCurrentUser(this *beego.Controller) models.User {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userId := sess.Get("userId")
	o := orm.NewOrm()
	o.Using("default")
	var user models.User

	if err := o.QueryTable(new(models.User)).Filter("id", userId).One(&user); err != nil {
		this.Redirect("/login", 302)
	}
	return user
}

func IsUserLogedIn(this *beego.Controller) bool {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userId := sess.Get("userId")
	return userId != nil
}

func GetORM() orm.Ormer  {
	o := orm.NewOrm()
	o.Using("default")
	return o
}
