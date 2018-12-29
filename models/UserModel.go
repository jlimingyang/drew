package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)

type UserBase struct {
	Id       int `orm:"pk;auto"`
	Username string
	Password string
	Token    string
	Status   int8
	Online   int8
	C_time   time.Time
}

//插入用户
func InsertUser(user *UserBase) (int64, error) {
	data := []byte(user.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	user.Password = md5str1
	log.Debug("密码:{}", user.Password)
	o := orm.NewOrm()
	o.Using("default")
	return o.Insert(user)
}
