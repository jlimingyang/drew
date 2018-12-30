package test

import (
	"github.com/astaxie/beego"
	"item1024.com/draw/utils"
	"testing"
)

func TestGenToken(t *testing.T) {
	token := utils.GenToken(12, "测试", 1)
	beego.Debug("generate token is ", token)
	custom := utils.CheckToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMiIsInVzZXJOYW1lIjoi5rWL6K-VIiwiZXhwIjoxNTQ2MTQyNDU1LCJuYmYiOjE1NDYxNDE0NTV9.r3urir4itjWs6Xb2tQaJk-rM0p7jYoDhnJXtxIqg4yo ")
	beego.Debug(custom == nil)
	beego.Debug("token userId is ", custom.UserId)
	beego.Debug("token userName is ", custom.UserName)
}
