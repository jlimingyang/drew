package controllers

import (
	"github.com/astaxie/beego"
	"item1024.com/draw/constants"
	m "item1024.com/draw/models"
	"item1024.com/draw/utils"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
}

func (this *UserController) UserRegiste() {
	user := new(m.UserBase)
	user.Online = 0
	user.Username = this.GetString("username")
	user.Password = this.GetString("password")
	user.Token = this.GetString("usb")
	this.Data["username"] = user.Username
	this.Data["password"] = user.Password
	//判断参数
	if isEmpty(user.Username) || isEmpty(user.Password) || isEmpty(user.Token) {
		this.Data["msg"] = constants.MSG_PARAM_NULL
		this.TplName = "user/registe.tpl"
		return
	}
	//判断是否存在
	num := m.UserNameIsExist(user.Username)
	beego.Debug("用户是否存在:", num)
	if num {
		this.Data["msg"] = constants.MSG_REGISTE_ALREAD
		this.TplName = "user/registe.tpl"
		return
	}
	//判断邀请码是否存在
	token := m.TokenIsExist(user.Token)
	beego.Debug("token是否存在:", token)
	if !token {
		beego.Debug("token验证未通过")
		this.Data["msg"] = constants.MSG_TOKEN_INVAIL
		this.TplName = "user/registe.tpl"
		return
	}
	user.C_time = time.Now()
	user.Status = 1
	error := m.InsertUser(user)
	if error != nil {
		beego.Debug(error)
		this.Data["msg"] = constants.MSG_REGIEST_FAIL
		this.TplName = "user/registe.tpl"
		return
	} else {
		this.Data["msg"] = constants.MSG_REGIEST_SUC
		this.Data["msgType"] = 1
		this.TplName = "user/login.tpl"
		return
	}
}

func (this *UserController) UserLogin() {
	username := this.GetString("username")
	password := this.GetString("password")
	this.Data["username"] = username
	this.Data["password"] = password
	//判断参数
	if isEmpty(username) || isEmpty(password) {
		this.Data["msg"] = constants.MSG_PARAM_NULL
		this.TplName = "user/login.tpl"
		return
	}
	if m.UserNameAndPasswordIsExist(username, password) {
		//查询user
		userBase := m.QueryUserByUserName(username)
		beego.Debug("查询到的用户名:", userBase.Username)
		ext, _ := strconv.ParseInt(beego.AppConfig.String("tokenExt"), 10, 64)
		//生成token
		token := utils.GenToken(userBase.Id, userBase.Username, ext)
		beego.Debug("auth:", token)
		this.Data["token"] = token
		this.Data["username"] = username
		this.TplName = "admin/index.html"
	} else {
		this.Data["msg"] = constants.MSG_LOGIN_FAIL
		this.TplName = "user/login.tpl"
		return
	}
}

func (this *UserController) Exit() {
	userId, _ := this.GetInt("userId")
	m.UpdateUserOnline(userId, 0)
	this.Data["msg"] = "debug"
	this.TplName = "user/login.tpl"
	return
}
