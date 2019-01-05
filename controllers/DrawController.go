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
	var res int64 = -1
	if num > 10000 {
		this.Data["json"] = map[string]interface{}{"code": 201, "success": true, "message": "抽奖人数不能大于1w"}
	} else {
		res = models.SaveDraw(userId, num)
	}
	if res != -1 {
		this.Data["json"] = map[string]interface{}{"code": 200, "success": true, "message": "设置成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 201, "success": false, "message": "设置失败"}
	}
	this.ServeJSON()
	return
}

func (this *DrawController) QueryDraw() {
	userId, _ := this.GetInt("userId")
	page, _ := this.GetInt("page")
	pagesize, _ := this.GetInt("pagesize")
	beego.Debug("userId:", userId, " - page:", page, " - pagesize:", pagesize)
	list := models.QueryDrawByUserId(page, pagesize, userId)
	this.Data["json"] = map[string]interface{}{"code": 200, "success": true, "message": "suc", "data": list}
	this.ServeJSON()
	return
}

func (this *DrawController) QueryDrawRecord() {
	userId, _ := this.GetInt("userId")
	level, _ := this.GetInt("level")
	page, _ := this.GetInt("page")
	pagesize, _ := this.GetInt("pagesize")
	beego.Debug("userId:", userId, " - page:", page, " - pagesize:", pagesize)
	list := models.QueryDrawRecordByUserId(page, pagesize, userId, level)
	this.Data["json"] = map[string]interface{}{"code": 200, "success": true, "message": "suc", "data": list}
	this.ServeJSON()
	return
}

func (this *DrawController) UpdateNameById() {
	name := this.GetString("name")
	id, _ := this.GetInt("id")
	_, err := models.UpdateDrawNameById(id, name)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 200, "success": true, "message": "更新成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 201, "success": false, "message": "更新失败"}
	}
	this.ServeJSON()
	return
}

func (this *DrawController) DeleteById() {
	id, _ := this.GetInt("id")
	if models.DeleteDrawById(id) {
		this.Data["json"] = map[string]interface{}{"code": 200, "success": true, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 201, "success": false, "message": "删除失败"}
	}
	this.ServeJSON()
	return
}

//保存抽奖
func (this *DrawController) SaveDraw() {
	row, _ := this.GetInt("row")
	userId, _ := this.GetInt("userId")
	level, _ := this.GetInt("level")
	_, err := models.SaveDrawRecord(userId, row, level)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 200, "success": true, "message": "保存成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 201, "success": false, "message": "保存失败"}
	}
	this.ServeJSON()
	return
}
