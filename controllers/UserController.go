package controllers

import (
	"github.com/astaxie/beego"
	m "item1024.com/draw/models"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) UserLogin() {
	user := new(m.TUserBase)
	user.Online = 1
	user.Username = u.GetString("userName")
	user.Password = u.GetString("userPwd")
	user.C_time = time.Now().Format("2006-01-02 15:04:05")
	user.Status = 1
	res, error := m.User_insert(user)
	println("插入成功:" + strconv.FormatInt(res, 10))
	if error != nil {
		println(error)
	}
}
