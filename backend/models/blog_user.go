package models

type BlogUser struct {
	Id           int    `xorm:"not null pk autoincr INT(10)"`
	Username     string `xorm:"not null VARCHAR(16)"`
	Password     string `xorm:"not null VARCHAR(255)"`
	Avatar       string `xorm:"VARCHAR(255)"`
	Introduction string `xorm:"VARCHAR(255)"`
	Nickname     string `xorm:"VARCHAR(32)"`
	About        string `xorm:"TEXT"`
}
