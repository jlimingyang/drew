package models

import (
	"github.com/astaxie/beego/orm"
	"item1024.com/draw/utils"
	"time"
)

type UserToken struct {
	Id     int `orm:"pk;auto"`
	UserId int
	Token  string `orm:"unique"`
	Status int8   "0为无效  1为有效  2为已经使用"
	C_time time.Time
	U_time time.Time
}

type User struct {
	userId   int    `json:"Id"`
	userName string `json:"Issuer"`
}

//插入token
func InsertToken(token *UserToken) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	token.C_time = time.Now()
	token.U_time = time.Now()
	token.Status = 1
	return o.Insert(token)
}

//设置token状态 0 无效  1 有效 2 已使用
func setInvalidById(id int, status int8) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	userToken := new(UserToken)
	userToken.Status = status
	return o.Update(&userToken, "Status")
}

// 查询token列表
func QueryToken(page int, pagesize int) (list utils.Page) {
	page -= 1
	if page-1 < 0 {
		page = 0
	}
	pagesize = page * pagesize
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(UserToken))
	qs.OrderBy("-id")
	qs.Limit(page, pagesize)
	total, error := qs.All(&list)
	if error != nil {
		utils.PageUtil(int(total), page, pagesize, list)
	}
	return
}

//查询
func TokenIsExist(token string) bool {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(UserToken))
	qs = qs.Filter("Token", token).Filter("Status", "1")
	return qs.Exist()
}
