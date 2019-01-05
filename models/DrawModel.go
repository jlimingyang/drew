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
	//检验是否已经存在
	if DrawConfigIsExist(userId) {
		//清空相关数据
		if !DeleteDrawByUserId(userId) {
			return -1
		}
	}
	//清除中奖名单
	if DrawRecordIsExist(userId) {
		//清空相关数据
		if !DeleteRecordByUserId(userId) {
			return -1
		}
	}
	draw := make([]DrawConfig, num)
	for i := 0; i < num; i++ {
		draw[i].Row = i + 1
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

func DrawConfigIsExist(userId int) bool {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawConfig)).Filter("UserId", userId)
	return qs.Exist()
}

func DrawRecordIsExist(userId int) bool {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawRecord)).Filter("UserId", userId)
	return qs.Exist()
}

func DeleteDrawByUserId(userId int) bool {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawConfig)).Filter("UserId", userId)
	_, err := qs.Delete()
	if err == nil {
		return true
	}
	return false
}

func DeleteRecordByUserId(userId int) bool {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawRecord)).Filter("UserId", userId)
	_, err := qs.Delete()
	if err == nil {
		return true
	}
	return false
}

func DeleteDrawById(id int) bool {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawConfig)).Filter("Id", id)
	_, err := qs.Delete()
	if err == nil {
		return true
	}
	return false
}

func QueryDrawByUserId(page int, size int, userId int) utils.Page {
	page -= 1
	if page-1 < 0 {
		page = 0
	}
	var list []*DrawConfig
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawConfig)).Filter("UserId", userId).OrderBy("-id").Limit(size, page*size)
	total, _ := qs.Count()
	_, error := qs.All(&list, "Id", "Name", "Row", "C_time", "U_time")
	if error == nil {
		return utils.PageUtil(int(total), page+1, size, list)
	}
	return utils.PageUtil(0, page+1, size, list)
}

func UpdateDrawNameById(id int, name string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	return o.QueryTable(new(DrawConfig)).Filter("Id", id).Update(orm.Params{
		"Name": name, "U_time": time.Now(),
	})
}

func QueryNameByRow(row int) string {
	o := orm.NewOrm()
	o.Using("default")
	var draw DrawConfig
	err := o.QueryTable(new(DrawConfig)).Filter("Row", row).One(&draw)
	if err == nil {
		return draw.Name
	}
	return ""
}

func SaveDrawRecord(userId int, row int, level int) (int64, error) {
	name := QueryNameByRow(row)
	drawRecord := DrawRecord{UserId: userId, Row: row, Name: name, Level: level, C_time: time.Now(), U_time: time.Now()}
	o := orm.NewOrm()
	o.Using("default")
	return o.Insert(&drawRecord)
}

func QueryDrawRecordByUserId(page int, size int, userId int, level int) utils.Page {
	page -= 1
	if page-1 < 0 {
		page = 0
	}
	var list []*DrawRecord
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(DrawRecord)).Filter("UserId", userId)
	if level > 0 {
		qs = qs.Filter("Level", level)
	}
	qs.OrderBy("-id").Limit(size, page*size)
	total, _ := qs.Count()
	_, error := qs.All(&list)
	if error == nil {
		utils.PageUtil(int(total), page, size, list)
	}
	return utils.PageUtil(0, page+1, size, list)
}
