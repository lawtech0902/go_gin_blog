package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/lawtech0902/go_gin_blog/backend/middleware/logger"
	"os"
)

/*
通用工具包
*/

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum([]byte(nil)))
}

func WriteErrorLog(s string) {
	file, _ := os.OpenFile(logger.GetLogFileFullPath("error"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	_, e := file.WriteString(s)
	if e != nil {
		fmt.Println(e)
	}
}
