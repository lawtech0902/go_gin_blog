package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/app"
	"github.com/lawtech0902/go_gin_blog/backend/service/article_service"
	"github.com/lawtech0902/go_gin_blog/backend/service/comment_service"
	"github.com/unknwon/com"
	"net/http"
)

func GetArticleComments(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	article := article_service.Article{ID: id}
	
	comments, err := article.GetCommentsByArticle()
	if err != nil {
		c.JSON(http.StatusNotFound, app.GenResponse(40020, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, comments, nil))
}

func CreateComment(c *gin.Context) {
	comment := comment_service.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	co, err := comment.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40024, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, co, nil))
}

func GetAllComments(c *gin.Context) {
	limit := c.DefaultQuery("limit", "")
	page := c.DefaultQuery("page", "")
	title := c.DefaultQuery("title", "")
	
	if title != "" {
		comments, err := comment_service.GetCommentsByArticleName(title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40032, nil, err))
			return
		}
		
		c.JSON(http.StatusOK, app.GenResponse(20000, comments, nil))
		return
	}
	
	comments, err := comment_service.Comment{}.GetAll(limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40032, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, comments, nil))
}

func DeleteComment(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	comment := comment_service.Comment{ID: uint(id)}
	
	co, err := comment.GetOne()
	if err != nil {
		c.JSON(http.StatusNotFound, app.GenResponse(40034, nil, err))
		return
	}
	
	if err := co.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40035, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, co, nil))
}
