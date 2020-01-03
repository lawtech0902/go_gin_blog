package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/app"
	"github.com/lawtech0902/go_gin_blog/backend/service/upload_service"
	"net/http"
)

func UploadImageAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	t := c.PostForm("t")
	if err != nil {
		c.JSON(http.StatusBadRequest, app.GenResponse(40000, nil, err))
		return
	}
	
	imageInfo, err := upload_service.UploadImageAvatarService(file, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40029, nil, err))
		return
	}
	
	dst := imageInfo.ImageRelPath
	if dst == "" {
		dst = imageInfo.AvatarRelPath
	}
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, app.GenResponse(40030, nil, err))
		return
	}
	
	c.JSON(http.StatusOK, app.GenResponse(20000, imageInfo, nil))
}
