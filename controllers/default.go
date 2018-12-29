package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	println("我就是想试一下打印的感觉")
	v := c.GetSession("test")
	if v == nil {
		c.SetSession("test", string("测试测试"))
	} else {
		print(v)
	}
	c.Data["Website"] = "测试控制器"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
