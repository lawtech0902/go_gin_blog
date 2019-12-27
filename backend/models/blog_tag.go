package models

type BlogTag struct {
	Id      int    `xorm:"not null pk autoincr INT(10)"`
	TagName string `xorm:"not null VARCHAR(16)"`
}
