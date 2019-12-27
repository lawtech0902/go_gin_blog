package models

type BlogArticle struct {
	Id          int    `xorm:"not null pk autoincr INT(10)"`
	Title       string `xorm:"not null unique VARCHAR(32)"`
	Content     string `xorm:"not null TEXT"`
	Html        string `xorm:"not null TEXT"`
	CategoryId  int    `xorm:"not null index INT(10)"`
	CreatedTime string `xorm:"not null VARCHAR(32)"`
	UpdatedTime string `xorm:"default '' VARCHAR(32)"`
	Status      string `xorm:"not null default 'published' VARCHAR(16)"`
}
