package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"beego-fileServer/models"
	"beego-fileServer/models/helpers"
	"github.com/astaxie/beego"
	"strings"
)

type UserController struct {
	beego.Controller
}


func (this *UserController) Register() {
	helpers.SetLayoutFor(&this.Controller)
	this.TplName = "forms/register.tpl"

	if helpers.IsUserLogedIn(&this.Controller){
		var user = helpers.GetCurrentUser(&this.Controller)
		this.Redirect("/user/"+user.Login, 302)
	}

	if this.Ctx.Input.Method() == "GET" {
		return
	}

	login := strings.Trim(this.GetString("login", ""), " ")
	pass := this.GetString("password", "")
	repass := this.GetString("repassword", "")

	if login == "" || pass == "" || repass == "" {
		this.Data["Error"] = "Incorect login or password"
		return
	}

	o := helpers.GetORM()
	login = strings.ToLower(login)
	exist := o.QueryTable(new(models.User)).Filter("login", login).Exist()
	if exist {
		this.Data["Error"] = "User exist"
		return
	}

	if pass == repass {

		pass = pass + models.GetSalt()
		hasher := md5.New()
		hasher.Write([]byte(pass))
		hashedPass := hex.EncodeToString(hasher.Sum(nil))
		user := models.User{Login:login, Password:hashedPass}

		if id, err := o.Insert(&user); err == nil {
			sess := this.StartSession()
			defer sess.SessionRelease(this.Ctx.ResponseWriter)
			sess.Set("userId", id)
			//this.Redirect("/user/" + login, 302)
			this.Redirect("/upload", 302)
		} else {
			this.Data["Error"] = err
		}

	}else {
		this.Data["Error"] = "Passwords missmatch"
	}
}

func (this *UserController) Login()  {
	helpers.SetLayoutFor(&this.Controller)
	this.TplName = "forms/login.tpl"

	if helpers.IsUserLogedIn(&this.Controller){
		var user = helpers.GetCurrentUser(&this.Controller)
		this.Redirect("/user/"+user.Login, 302)
	}

	if this.Ctx.Input.Method() == "GET" {
		return
	}



	login := strings.Trim(this.GetString("login", ""), " ")
	pass := this.GetString("password", "")

	if login == "" || pass == "" {
		this.Data["Error"] = "Incorect login or password"
	}

	login = strings.ToLower(login)
	pass = pass + models.GetSalt()
	hasher := md5.New()
	hasher.Write([]byte(pass))
	hashedPass := hex.EncodeToString(hasher.Sum(nil))

	var o = helpers.GetORM()

	var user models.User
	err := o.QueryTable(new(models.User)).Filter("login", login).Filter("password", hashedPass).One(&user)

	if err != nil {
		this.Data["Error"] = "User not found"
	} else {
		sess := this.StartSession()
		defer sess.SessionRelease(this.Ctx.ResponseWriter)
		sess.Set("userId", user.Id)
		this.Redirect("/user/" + user.Login, 302)
	}
}

func (this *UserController) Logout() {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	sess.Delete("userId")
	sess.SessionRelease(this.Ctx.ResponseWriter)

	this.Redirect("/", 302)
}

