package routers

import (
	"github.com/astaxie/beego"
	"item1024.com/draw/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/registe.py", &controllers.UserController{}, "get:UserRegiste")
	//beego.Router("/login", &controllers.UserController{}, "get:UserLogin")

}
