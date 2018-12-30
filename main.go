package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"item1024.com/draw/constants"
	_ "item1024.com/draw/routers"
	"item1024.com/draw/utils"
	"strings"
	"time"
)

func init() {
	//连接数据库
	database := beego.AppConfig.String("mysqldb")
	dburl := beego.AppConfig.String("mysqlurls") + ":" + beego.AppConfig.String("mysqlport")
	dbalias := beego.AppConfig.String("mysqlaliasname")
	dbdrivername := beego.AppConfig.String("mysqldrivername")
	dbuser := beego.AppConfig.String("mysqluser")
	dbapassword := beego.AppConfig.String("mysqlpass")
	sqlDebug, error := beego.AppConfig.Bool("sqldebug")
	println(error)
	println("runmode:" + beego.AppConfig.String("runmode") + "--connect db:" + database + "--url:" + dburl)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbalias, dbdrivername, dbuser+":"+dbapassword+"@tcp("+dburl+")/"+database+"?charset=utf8")
	orm.Debug = sqlDebug
	// create table
	orm.RunSyncdb(dbalias, false, true)
	orm.DefaultTimeLoc = time.UTC

}

func main() {
	beego.Run()
	beego.SetLogger("log", `{"filename":"logs/draw.log"}`)
	beego.BeeLogger.DelLogger("log")
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	beego.InsertFilter("*/", beego.BeforeExec, FilterUser)
}

var FilterUser = func(ctx *context.Context) {
	res := utils.Response{500, constants.MSG_LOGIN_FAIL, nil}
	msg, err := json.Marshal(res)
	token := ctx.Request.Header.Get("auth")
	if !strings.Contains(token, ".") {
		if err == nil {
			custom := utils.CheckToken(token)
			if custom != nil {
				ctx.Input.SetParam("userId", custom.UserId)
				ctx.Input.SetParam("userName", custom.UserName)
			}
		}
	}
	ctx.ResponseWriter.Write(msg)

}
