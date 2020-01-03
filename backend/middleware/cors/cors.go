package cors

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var (
			// filterHost 做过滤器，防止不合法的域名访问
			filterHost = [...]string{"http://localhost.*"}
			isAccess   = false
		)
		
		for _, v := range filterHost {
			match, _ := regexp.MatchString(v, origin)
			if match {
				isAccess = true
			}
		}
		
		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Token")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		
		// 履行所有Options方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		
		c.Next()
	}
}
