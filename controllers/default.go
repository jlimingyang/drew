package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	println(c.GetString("userId"))
	println(c.Ctx.Request.Header.Get("userId"))
	v := c.GetSession("test")
	if v == nil {
		c.SetSession("test", string("测试测试"))
	} else {
		print(v)
	}
	beego.Debug(test(555555), "1232")
	c.Data["Website"] = "测试控制器"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func test(i int) (str string) {
	str = string(i) + "hello"
	return
}
