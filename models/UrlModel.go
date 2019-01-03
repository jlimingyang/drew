package models

import "time"

type UrlBase struct {
	Id       int `orm:"pk;auto"`
	UserId   int
	ShortUrl string `orm:"unique"`
	LongUrl  string
	C_time   time.Time
	U_time   time.Time
}
