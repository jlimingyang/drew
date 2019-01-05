package controllers

import "github.com/astaxie/beego"

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

//抽奖设置页面
func (this *MainController) DrawSet() {
	token := this.GetSession("auth")
	beego.Debug("token:", token)
	this.Data["token"] = token
	this.TplName = "admin/draw_set.html"
}

//中奖奖页面
func (this *MainController) DrawList() {
	token := this.GetSession("auth")
	beego.Debug("token:", token)
	this.Data["token"] = token
	this.TplName = "admin/draw_list.html"
}

//抽奖页面
func (this *MainController) Draw() {
	token := this.GetSession("auth")
	beego.Debug("token:", token)
	this.Data["token"] = token
	this.TplName = "draw/index.html"
}
