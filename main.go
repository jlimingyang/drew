package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "item1024.com/draw/routers"
)

func init() {
	//连接数据库
	database := beego.AppConfig.String("mysqldb")
	dburl := beego.AppConfig.String("mysqlurls") + ":" + beego.AppConfig.String("mysqlport")
	dbalias := beego.AppConfig.String("mysqlaliasname")
	dbdrivername := beego.AppConfig.String("mysqldrivername")
	dbuser := beego.AppConfig.String("mysqluser")
	dbapassword := beego.AppConfig.String("mysqlpass")
	println("runmode:" + beego.AppConfig.String("runmode") + "--connect db:" + database + "--url:" + dburl)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbalias, dbdrivername, dbuser+":"+dbapassword+"@tcp("+dburl+")/"+database+"?charset=utf8")
}

func main() {
	beego.Run()
}
