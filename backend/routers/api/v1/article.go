package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/app"
	"github.com/lawtech0902/go_gin_blog/backend/service/article_service"
	"github.com/unknwon/com"
	"net/http"
)

func GetArticles(c *gin.Context) {
	limit := c.DefaultQuery("limit", "")
	page := c.DefaultQuery("page", "")
	category := c.DefaultQuery("category", "")
	q := c.DefaultQuery("q", "")
	tag := c.DefaultQuery("tag", "")
	key := c.DefaultQuery("key", "")
	status := c.DefaultQuery("status", "")
	admin := c.DefaultQuery("admin", "")
	
	if category != "" {
		data, err := article_service.GetArticlesByCategory(
			article_service.SetLimitPage(limit, page),
			article_service.SetAdmin(admin),
			article_service.SetCategory(category))
		if err != nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40022, nil, err))
			return
		}
		
		c.JSON(http.StatusOK, app.GenResponse(20000, data, nil))
		return
	}
	
	if tag != "" {
		data, err := article_service.GetArticlesByTag(
			article_service.SetLimitPage(limit, page),
			article_service.SetAdmin(admin),
			article_service.SetTag(tag))
		if err != nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40022, nil, err))
			return
		}
		
		c.JSON(http.StatusOK, app.GenResponse(20000, data, nil))
		return
	}
	
	if key != "" || status == "" {
		data, err := article_service.SearchArticle(key, status,
			article_service.SetLimitPage(limit, page),
			article_service.SetAdmin(admin),
			article_service.SetSearch(true))
		if err != nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40022, nil, err))
			return
		}
		
		c.JSON(http.StatusOK, app.GenResponse(20000, data, nil))
		return
	}
	
	if q != "" {
		data, err := article_service.SearchFromES(article_service.SetQ(q), article_service.SetLimitPage(limit, page))
		if err != nil {
			c.JSON(http.StatusInternalServerError, app.GenResponse(40022, nil, err))
			return
		}
		
		c.JSON(http.StatusOK, app.GenResponse(20000, data, nil))
		return
	}
	
	article := article_service.Article{}
	data, err := article.GetAll(article_service.SetLimitPage(limit, page), article_service.SetAdmin(admin))
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40022, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, data, nil))
}

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	admin := c.DefaultQuery("admin", "")
	article := article_service.Article{ID: id}
	
	articleDetail, err := article.GetOne(article_service.SetAdmin(admin))
	if err != nil {
		c.JSON(http.StatusNotFound, app.GenResponse(40020, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, articleDetail, nil))
}

func CreateArticle(c *gin.Context) {
	article := article_service.Article{}
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40024, nil, err))
	}
	
	a, err := article.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40024, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, a, nil))
}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	admin := c.DefaultQuery("admin", "")
	article := article_service.Article{ID: id}
	
	articleDetail, _ := article.GetOne(article_service.SetAdmin(admin))
	a := articleDetail.A
	
	if err := a.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40026, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, a, nil))
}

func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	admin := c.DefaultQuery("admin", "")
	article := article_service.Article{ID:id}
	
	articleDetail, _ := article.GetOne(article_service.SetAdmin(admin))
	a := articleDetail.A
	
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40024, nil, err))
		return
	}
	
	if err := a.Edit(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40025, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, a, nil))
}
