package upload_service

import (
	"github.com/lawtech0902/go_gin_blog/backend/pkg/upload"
	"mime/multipart"
	"path/filepath"
)

func UploadImageAvatarService(file *multipart.FileHeader, t string) (info upload.ImageAvatarInfo, err error) {
	name := filepath.Base(file.Filename)
	filename := upload.MD5ImageName(name)
	
	info = upload.GenImageAvatarInfo(t, filename)
	relPath := filepath.Dir(info.AvatarRelPath)
	if relPath == "." {
		relPath = filepath.Dir(info.ImageRelPath)
	}
	
	if err = upload.IsNotExistMkDir(relPath); err != nil {
		return
	}
	
	return
}
