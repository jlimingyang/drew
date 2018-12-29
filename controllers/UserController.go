package controllers

import (
	"github.com/astaxie/beego"
	"item1024.com/draw/constants"
	m "item1024.com/draw/models"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
}

func (u *UserController) UserRegiste() {
	user := new(m.UserBase)
	user.Online = 0
	user.Username = u.GetString("userName")
	user.Password = u.GetString("userPwd")
	user.C_time = time.Now()
	user.Status = 1
	res, error := m.InsertUser(user)
	beego.Debug("insert sucs:", strconv.FormatInt(res, 10))
	if error != nil {
		println(error)
	} else {
		u.Data["msg"] = constants.MSG_LOGIN_SUC
		u.TplName = "index.tpl"
	}
}
