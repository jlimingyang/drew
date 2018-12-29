package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(TUserBase))
}

type TUserBase struct {
	Id       int
	Username string
	Password string
	Status   int8
	Online   int8
	C_time   string
}

func User_insert(user *TUserBase) (int64, error) {

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	return o.Insert(user)
}
