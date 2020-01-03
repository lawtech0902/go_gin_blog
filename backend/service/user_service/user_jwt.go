package user_service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/utils"
	"time"
)

/*
用户jwt身份鉴权
*/

var jwtSecret = []byte(setting.AppInfo.JwtSecret)

type CustomClaims struct {
	User
	jwt.StandardClaims
}

func (u User) CheckAuth() bool {
	var count int
	if err := db.Get(&count, "select count(1) from blog_user where username=? and password=?", u.Username, utils.EncodeMD5(u.Password)); err != nil {
		utils.WriteErrorLog(err.Error())
		return false
	}
	
	return count > 0
}

func (u User) GenToken() (string, error) {
	claims := CustomClaims{u, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(setting.AppInfo.TokenTimeout)).Unix(),
		Id:        fmt.Sprintf("%v", time.Now().UnixNano()),
	}}
	
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	
	if token != nil {
		if _, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return nil
		} else {
			return err
		}
	}
	
	return err
}
