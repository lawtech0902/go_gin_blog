package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/app"
	"github.com/lawtech0902/go_gin_blog/backend/service/user_service"
	"net/http"
)

func Login(c *gin.Context) {
	data := make(map[string]interface{})
	user := user_service.User{}
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	isExist := user.CheckAuth()
	if isExist {
		token, err := user.GenToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40004, nil, err))
			return
		}
		data["token"] = token
		c.JSON(http.StatusOK, app.GenResponse(20000, data, nil))
		return
	}
	
	c.JSON(http.StatusUnauthorized, app.GenResponse(40001, nil, nil))
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, app.GenResponse(20000, nil, nil))
}

func GetUserInfo(c *gin.Context) {
	userInfo, err := user_service.GetUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40027, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, userInfo, nil))
}

func GetUserAbout(c *gin.Context) {
	about, err := user_service.GetAbout()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40027, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, about, nil))
}

func EditUser(c *gin.Context) {
	bytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40028, nil, err))
		return
	}
	
	u := user_service.User{}
	if err := json.Unmarshal(bytes, &u); err!= nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40028, nil, err))
		return
	}
	
	if u.Password != "" {
		if err := u.ResetPassword(); err!= nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40028, nil, err))
			return
		}
	} else if u.About != "" {
		if err := u.EditAbout(); err != nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40028, nil, err))
			return
		}
	} else {
		if err := u.EditUser(); err != nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40028, nil, err))
			return
		}
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, u, nil))
}
