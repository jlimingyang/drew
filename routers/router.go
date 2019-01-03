package routers

import (
	"github.com/astaxie/beego"
	"item1024.com/draw/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Login")
	beego.Router("/login.page", &controllers.MainController{}, "get:Login")
	beego.Router("/registe.page", &controllers.MainController{}, "get:Registe")

	beego.Router("/token/addToken.py", &controllers.TokenController{}, "get:CreateToken")
	beego.Router("/user/registe.py", &controllers.UserController{}, "post:UserRegiste")
	beego.Router("/user/login.py", &controllers.UserController{}, "post:UserLogin")

	beego.Router("/auth/page/saveDrawConfig", &controllers.DrawController{}, "post:SaveDrawConfig")

}
