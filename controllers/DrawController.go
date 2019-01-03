package controllers

import (
	"github.com/astaxie/beego"
	"item1024.com/draw/models"
)

type DrawController struct {
	BaseController
}

func (this *DrawController) SaveDrawConfig() {
	userId, _ := this.GetInt("userId")
	beego.Debug("用户Id:", userId)
	num, _ := this.GetInt("num")
	models.SaveDraw(userId, num)
	this.TplName = ""
}

func (this *DrawController) QueryDraw() {
	userId, _ := this.GetInt("userId")
	page, _ := this.GetInt("page")
	pagesize, _ := this.GetInt("pagesize")
	list := models.QueryDrawByUserId(page, pagesize, userId)
	this.Data["json"] = map[string]interface{}{"code": 200, "success": true, "message": "suc", "data": list}
	this.ServeJSON()
	return
}
