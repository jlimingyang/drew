package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"item1024.com/draw/utils"
	"time"
)

type DrawConfig struct {
	Id     int `orm:"pk;auto"`
	UserId int
	Row    int
	Name   string
	C_time time.Time
	U_time time.Time
}

type DrawRecord struct {
	Id     int `orm:"pk;auto"`
	UserId int
	Row    int
	Name   string
	Level  int
	C_time time.Time
	U_time time.Time
}

func SaveDraw(userId int, num int) (count int64) {
	draw := make([]DrawConfig, num)
	for i := 1; i <= num; i++ {
		draw[i].Row = i
		draw[i].C_time = time.Now()
		draw[i].UserId = userId
		draw[i].U_time = time.Now()
	}
	o := orm.NewOrm()
	o.Using("default")
	suc, err := o.InsertMulti(100, draw)
	if err != nil {
		beego.Debug(err)
	}
	return suc
}

func QueryDrawByUserId(page int, size int, userId int) (list utils.Page) {
	page -= 1
	if page-1 < 0 {
		page = 0
	}
	size = page * size
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawConfig)).Filter("UserId", userId).OrderBy("-id").Limit(page, size)
	total, error := qs.All(&list)
	if error == nil {
		utils.PageUtil(int(total), page, size, list)
	}
	return
}

func UpdateDrawNameById(id int, name string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	return o.QueryTable(new(DrawConfig)).Filter("Id", id).Update(orm.Params{
		"Name": name, "U_time": time.Now(),
	})
}

func saveDrawRecord(userId int, row int, name string, level int) (int64, error) {
	drawRecord := DrawRecord{UserId: userId, Row: row, Name: name, Level: level, C_time: time.Now(), U_time: time.Now()}
	o := orm.NewOrm()
	o.Using("default")
	return o.Insert(drawRecord)
}

func QueryDrawRecordByUserId(page int, size int, userId int) (list utils.Page) {
	page -= 1
	if page-1 < 0 {
		page = 0
	}
	size = page * size
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawRecord)).Filter("UserId", userId).OrderBy("-id").Limit(page, size)
	total, error := qs.All(&list)
	if error == nil {
		utils.PageUtil(int(total), page, size, list)
	}
	return
}
