package models

import (
	"github.com/astaxie/beego/orm"
	"item1024.com/draw/utils"
	"time"
)

type UserToken struct {
	Id     int `orm:"pk;auto"`
	UserId int
	Token  string
	Status int8
	C_time time.Time
	U_time time.Time
}

//插入token
func InsertToken(token *UserToken) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	token.Status = 1
	return o.Insert(token)
}

// 查询token列表
func QueryToken(page int, pagesize int) (list utils.Page) {
	page -= 1
	if page-1 < 0 {
		page = 0
	}
	pagesize = page * pagesize
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserToken))
	qs.OrderBy("-id")
	qs.Limit(page, pagesize)
	total, error := qs.All(&list)
	if error != nil {
		utils.PageUtil(int(total), page, pagesize, list)
	}
	return
}

//查找列表
func CountTokenByUserId(userId int64, status int8) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserToken))
	qs.Filter("userId", userId)
	qs.Filter("status", 1)
	return qs.Count()
}

//查询用户
func QueryUserTokenByToken(token string) (userToken UserToken) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserToken))
	qs.Filter("token", token)
	qs.Filter("status", 1)
	qs.One(userToken)
	return
}
