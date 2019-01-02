package controllers

import (
	"github.com/astaxie/beego"
	"item1024.com/draw/models"
)

type TokenController struct {
	BaseController
}

func (this *TokenController) CreateToken() {
	token := GenToken()
	beego.Debug("token:", token)
	//检查token是否存在
	if models.TokenIsExist(token) {
		this.Data["json"] = map[string]interface{}{"code": 503, "success": false, "message": "授权码重复"}
	}
	userToken := models.UserToken{Token: token}
	_, err := models.InsertToken(&userToken)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 200, "success": true, "message": "suc"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 502, "success": false, "message": "fail"}
	}
	this.ServeJSON()
	return
}
