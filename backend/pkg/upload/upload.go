package upload

import (
	"fmt"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageAvatarInfo struct {
	ImageFullUrl  string
	ImageRelPath  string
	AvatarFullUrl string
	AvatarRelPath string
}

func MD5ImageName(name string) string {
	ext := filepath.Ext(name)
	filename := utils.EncodeMD5(name + fmt.Sprintf("%v", time.Now().Unix()))
	return filename + ext
}

func GenImageAvatarInfo(t, n string) (i ImageAvatarInfo) {
	if t == "avatar" {
		i.ImageRelPath = ""
		i.ImageFullUrl = ""
		i.AvatarRelPath = filepath.Join(setting.AppInfo.RootBasePath, setting.AppInfo.UploadBasePath, setting.AppInfo.AvatarRelPath, n)
		i.AvatarFullUrl = fmt.Sprintf(`%s/%s/%s/%s`, setting.AppInfo.ApiBaseUrl, setting.AppInfo.AvatarRelPath, n)
	} else if t == "image" {
		i.ImageRelPath = filepath.Join(setting.AppInfo.UploadBasePath, setting.AppInfo.ImageRelPath, today(), n)
		i.ImageFullUrl = fmt.Sprintf(`%s/%s/%s/%s/%s`, setting.AppInfo.ApiBaseUrl, setting.AppInfo.ImageRelPath, today(), n)
		i.AvatarRelPath = ""
		i.AvatarFullUrl = ""
	}
	return
}

func today() string {
	t := time.Now().Format(strings.Split(setting.AppInfo.TimeFormat, " ")[0])
	return t
}

// 如果不存在返回true，存在则返回false
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	
	return os.IsNotExist(err)
}

func IsNotExistMkDir(src string) error {
	if CheckExist(src) {
		e := MkDir(src)
		return e
	}
	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModeDir)
	return err
}
