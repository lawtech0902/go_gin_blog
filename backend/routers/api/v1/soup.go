package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/app"
	"github.com/lawtech0902/go_gin_blog/backend/service/soup_service"
	"github.com/unknwon/com"
	"net/http"
)

func GetAllSoups(c *gin.Context) {
	limit := c.DefaultQuery("limit", "")
	page := c.DefaultQuery("page", "")
	soups, err := soup_service.Soup{}.GetAll(limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40036, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, soups, nil))
}

func CreateSoup(c *gin.Context) {
	soup := soup_service.Soup{}
	if err := c.ShouldBindJSON(&soup); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	s, err := soup.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40037, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, s, err))
}

func DeleteSoup(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	soup := soup_service.Soup{ID: id}
	s, err := soup.GetOne()
	if err != nil {
		c.JSON(http.StatusNotFound, app.GenResponse(40038, nil, err))
		return
	}
	
	if err = s.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40039, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, s, nil))
}

func EditSoup(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	soup := soup_service.Soup{ID: id}
	if err := c.ShouldBindJSON(&soup); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	if err := soup.Edit(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40040, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, soup, nil))
}

func GetRandSoup(c *gin.Context) {
	soup, err := soup_service.Soup{}.GetRandOne()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40041, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, soup, nil))
}
