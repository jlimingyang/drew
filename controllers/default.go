package controllers

type MainController struct {
	BaseController
}

//登录页面
func (this *MainController) Login() {
	this.Data["msg"] = "debug"
	this.TplName = "user/login.tpl"
}

//注册页面
func (this *MainController) Registe() {
	this.Data["msg"] = "debug"
	this.TplName = "user/registe.tpl"
}
