package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserBase struct {
	Id       int    `orm:"pk;auto"`
	Username string `orm:"unique"`
	Password string
	Token    string
	Status   int8
	Online   int8
	C_time   time.Time
}

//插入用户
func InsertUser(user *UserBase) error {
	data := []byte(user.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	user.Password = md5str1
	beego.Debug("密码:{}", user.Password)
	o := orm.NewOrm()
	o.Using("default")
	o.Begin()
	_, err := o.Insert(user)
	_, err02 := o.QueryTable(new(UserToken)).Filter("Token", user.Token).Update(orm.Params{
		"Status": "2", "UserId": user.Id,
	})
	if err != nil || err02 != nil {
		beego.Debug(err)
		beego.Debug(err02)
		return o.Rollback()
	}
	return o.Commit()
}

//username是否存在
func UserNameIsExist(userName string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserBase))
	qs = qs.Filter("Username", userName)
	return qs.Exist()
}

//username&password是否存在
func UserNameAndPasswordIsExist(usernName string, password string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserBase))
	has := md5.Sum([]byte(password))
	password = fmt.Sprintf("%x", has)
	qs = qs.Filter("UserName", usernName).Filter("Password", password).Filter("Status", 1)
	return qs.Exist()
}

//更改用户状态
func UpdateUserOnline(userId int, online int) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable(new(UserBase)).Filter("Id", userId).Update(orm.Params{
		"Online": online,
	})
	return err
}

//查询用户
func QueryUserByUserName(username string) (userbase UserBase) {
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(new(UserBase)).Filter("Username", username).One(&userbase, "Id", "Username")
	return
}
