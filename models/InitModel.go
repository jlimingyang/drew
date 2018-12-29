package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
初始化model
*/

func init() {
	// 需要在 init 中注册定义的 model
	orm.RegisterModelWithPrefix(beego.AppConfig.String("tbprefix"), new(UserToken), new(UserBase))
}
