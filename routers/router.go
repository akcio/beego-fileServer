package routers

import (
	"fileServer/testBeeGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	//beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitepoint")
	//
    //beego.Router("/list", &controllers.NewController{})
    beego.Router("/", &controllers.FilesController{}, "get,post:Upload")

    beego.Router("/register", &controllers.UserController{}, "get,post:Register")
    beego.Router("/login", &controllers.UserController{}, "get,post:Login")
    beego.Router("/logout", &controllers.UserController{}, "get:Logout")

    beego.Router("/user/:user:string", &controllers.FilesController{}, "get:List")

    beego.Router("/upload", &controllers.FilesController{}, "get,post:Upload")
    beego.Router("/download/:name([A-Za-z0-9\\:\\+_-]+)", &controllers.FilesController{}, "get,post:Download")
	beego.Router("/delete/:name([A-Za-z0-9\\:\\+_-]+)", &controllers.FilesController{}, "get,post:DeleteLink")
}
