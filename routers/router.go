package routers

import (
	"github.com/astaxie/beego"
	"item1024.com/draw/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Login")
	beego.Router("/login.page", &controllers.MainController{}, "get:Login")
	beego.Router("/registe.page", &controllers.MainController{}, "get:Registe")
	beego.Router("/draw_set.page", &controllers.MainController{}, "get:DrawSet")
	beego.Router("/draw_list.page", &controllers.MainController{}, "get:DrawList")
	beego.Router("/draw.page", &controllers.MainController{}, "get:Draw")

	beego.Router("/token/addToken.py", &controllers.TokenController{}, "get:CreateToken")
	beego.Router("/user/registe.py", &controllers.UserController{}, "post:UserRegiste")
	beego.Router("/user/login.py", &controllers.UserController{}, "post:UserLogin")
	beego.Router("/user/exit.py", &controllers.UserController{}, "get:Exit")

	beego.Router("/auth/api/saveDrawConfig", &controllers.DrawController{}, "post:SaveDrawConfig")
	beego.Router("/auth/api/queryDraw", &controllers.DrawController{}, "post:QueryDraw")
	beego.Router("/auth/api/updateNameById", &controllers.DrawController{}, "post:UpdateNameById")
	beego.Router("/auth/api/deleteById", &controllers.DrawController{}, "post:DeleteById")
	beego.Router("/auth/api/saveDraw", &controllers.DrawController{}, "post:SaveDraw")
	beego.Router("/auth/api/queryDrawRecord", &controllers.DrawController{}, "post:QueryDrawRecord")

}
