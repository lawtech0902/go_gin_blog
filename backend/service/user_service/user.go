package user_service

import (
	"github.com/lawtech0902/go_gin_blog/backend/models"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/utils"
)

type User struct {
	ID           int    `json:"id" db:"id" form:"id"`
	Username     string `json:"username" db:"username" form:"username"`
	Password     string `json:"password" db:"password" form:"password"`
	Introduction string `json:"introduction" db:"introduction" form:"introduction"`
	Avatar       string `json:"avatar" db:"avatar" form:"avatar"`
	Nickname     string `json:"nickname" db:"nickname" form:"nickname"`
	About        string `json:"about" db:"about" form:"about"`
}

var db = models.DB

// user crud 操作实现
func GetUser() (User, error) {
	var user User
	
	err := db.Get(&user, "select avatar, introduction, nickname from blog_user")
	return user, err
}

func (u *User) EditUser() error {
	_, err := db.Exec("update blog_user set introduction=?, avatar=?, nickname=?", u.Introduction, u.Avatar, u.Nickname)
	return err
}

func (u *User) ResetPassword() error {
	_, err := db.Exec("update blog_user set password=?", utils.EncodeMD5(u.Password))
	return err
}

func (u *User) EditAbout() error {
	_, err := db.Exec("update blog_user set about=?", u.About)
	return err
}

func GetAbout() (string, error) {
	var a string
	err := db.Get(&a, "select about from blog_user")
	return a, err
}
