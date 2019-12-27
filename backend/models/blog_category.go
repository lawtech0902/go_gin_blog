package models

type BlogCategory struct {
	Id           int    `xorm:"not null pk autoincr INT(10)"`
	CategoryName string `xorm:"VARCHAR(16)"`
}
