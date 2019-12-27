package models

type BlogComment struct {
	Id          int    `xorm:"not null pk autoincr INT(10)"`
	Username    string `xorm:"not null VARCHAR(16)"`
	IsAuthor    int    `xorm:"not null default 0 TINYINT(1)"`
	ParentId    int    `xorm:"index INT(10)"`
	RootId      int    `xorm:"index INT(10)"`
	ArticleId   int    `xorm:"not null index INT(10)"`
	Content     string `xorm:"not null VARCHAR(255)"`
	CreatedTime string `xorm:"not null VARCHAR(255)"`
}
