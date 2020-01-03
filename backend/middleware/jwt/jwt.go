package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/app"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/e"
	"github.com/lawtech0902/go_gin_blog/backend/service/user_service"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.Success
		var err error
		
		token := c.Query("token")
		if token == "" {
			code = e.TokenCheckError
		} else {
			err = user_service.ParseToken(token)
			if err != nil {
				code = e.TokenCheckError
			}
		}
		
		if code != e.Success {
			c.JSON(http.StatusUnauthorized, app.GenResponse(code, nil, err))
			c.Abort()
			return
		}
		
		c.Next()
	}
}
