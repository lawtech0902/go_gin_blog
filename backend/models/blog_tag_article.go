package models

type BlogTagArticle struct {
	Id        int `xorm:"not null pk autoincr INT(10)"`
	TagId     int `xorm:"not null index INT(10)"`
	ArticleId int `xorm:"not null index INT(10)"`
}
