package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	beego.Debug("goto")
	this.TplName = "user/login.tpl"
}

func test(i int) (str string) {
	str = string(i) + "hello"
	return
}
