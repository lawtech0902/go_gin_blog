package models

type BlogSoup struct {
	Id      int    `xorm:"not null pk autoincr INT(10)"`
	Content string `xorm:"not null TINYTEXT"`
}
