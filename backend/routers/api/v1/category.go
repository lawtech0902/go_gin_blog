package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/app"
	"github.com/lawtech0902/go_gin_blog/backend/service/category_service"
	"github.com/unknwon/com"
	"net/http"
)

func GetAllCategories(c *gin.Context) {
	categories, err := category_service.Category{}.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40015, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, categories, nil))
}

func CreateCategory(c *gin.Context) {
	category := category_service.Category{}
	if err := c.ShouldBindJSON(category); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	ca, err := category.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40016, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, ca, nil))
}

func DeleteCategory(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	category := category_service.Category{ID: id}
	cg, err := category.GetOne()
	if err != nil {
		c.JSON(http.StatusNotFound, app.GenResponse(40015, nil, err))
		return
	}
	
	if err = cg.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40018, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, cg, nil))
}

func EditCategory(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	category := category_service.Category{ID: id}
	
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	if err := category.Edit(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40017, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, category, nil))
}
