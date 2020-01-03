package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/app"
	"github.com/lawtech0902/go_gin_blog/backend/service/tag_service"
	"github.com/unknwon/com"
	"net/http"
)

func GetAllTags(c *gin.Context) {
	tags, err := tag_service.Tag{}.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40008, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, tags, nil))
}

func CreateTag(c *gin.Context) {
	tag := &tag_service.Tag{}
	if err := c.ShouldBindJSON(tag); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	t, err := tag.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40009, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, t, nil))
}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tag := tag_service.Tag{ID: id}
	
	t, err := tag.GetOne()
	if err != nil {
		c.JSON(http.StatusNotFound, app.GenResponse(40006, nil, err))
		return
	}
	
	if err := t.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40011, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, tag, nil))
}

func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tag := tag_service.Tag{ID: id}
	
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	if err := tag.Edit(); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40010, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, tag, nil))
}
