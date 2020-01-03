package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/middleware/cors"
	"github.com/lawtech0902/go_gin_blog/backend/middleware/jwt"
	"github.com/lawtech0902/go_gin_blog/backend/middleware/logger"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	v1 "github.com/lawtech0902/go_gin_blog/backend/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})
	r.Use(logger.CustomLogger(), cors.CorsMiddleware())
	r.Use(gin.Recovery())
	
	apiv1 := r.Group(setting.AppInfo.ApiBaseUrl)
	{
		apiv1.Static(setting.AppInfo.RootBasePath, setting.AppInfo.UploadBasePath)
		
		apiv1.POST("/user/login", v1.Login)
		apiv1.POST("/user/logout", v1.Logout)
		// GET方法和添加评论没有使用JWT认证
		apiv1.GET("/user/info", v1.GetUserInfo)
		apiv1.GET("/user/about", v1.GetUserAbout)
		apiv1.GET("/tags", v1.GetAllTags)
		apiv1.GET("/categories", v1.GetAllCategories)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.GET("/soup/random", v1.GetRandSoup)
		apiv1.GET("/comments/:id", v1.GetArticleComments)
		apiv1.POST("/comments", v1.CreateComment)
		
		apiv1.Use(jwt.JWT())
		apiv1.PATCH("/user/edit", v1.EditUser)
		
		apiv1.POST("/tags", v1.CreateTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		
		apiv1.POST("/categories", v1.CreateCategory)
		apiv1.PUT("/categories/:id", v1.EditCategory)
		apiv1.DELETE("/categories/:id", v1.DeleteCategory)
		
		apiv1.POST("/articles", v1.CreateArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		
		apiv1.POST("/soups", v1.CreateSoup)
		apiv1.DELETE("/soups/:id", v1.DeleteSoup)
		apiv1.PUT("/soups/:id", v1.EditSoup)
		apiv1.GET("/soups", v1.GetAllSoups)
		
		apiv1.GET("/comments", v1.GetAllComments)
		apiv1.DELETE("/comments/:id", v1.DeleteComment)
		
		apiv1.POST("/upload", v1.UploadImageAvatar)
	}
	
	return r
}
