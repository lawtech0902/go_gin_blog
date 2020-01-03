package article_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/gredis"
)

func (a Article) ViewKey() string {
	viewKey := a.Title + ":view"
	return viewKey
}

func addView(key string) error {
	err := gredis.INCRKey(key)
	return err
}

func getViews(key string) (n uint8, err error) {
	data, err := gredis.GetKey(key)
	if err != nil || data == nil {
		return n, err
	}
	
	v, ok := data.([]uint8)
	if ok {
		if err := json.Unmarshal([]byte(v[:]), &n); err != nil {
			return n, err
		}
		return n, nil
	} else {
		return n, errors.New("返回数据类型有误，json无法解析")
	}
}

func articleCacheKey(opts Options) string {
	if opts.Admin {
		return fmt.Sprintf("article_%d_%d_%s_%s_%s", opts.Limit, opts.Page, "admin", opts.C, opts.T)
	} else {
		return fmt.Sprintf("article_%d_%d_%s_%s", opts.Limit, opts.Page, opts.C, opts.T)
	}
}

func setArticleCache(key string, value Articles) error {
	marshal, _ := json.Marshal(value)
	err := gredis.SetKey(key, marshal, gredis.SetTimeout(true))
	return err
}

func getArticleCache(key string) (as Articles, err error) {
	data, err := gredis.GetKey(key)
	if err != nil || data == nil {
		return as, err
	}
	
	v, ok := data.([]uint8)
	if ok {
		if e := json.Unmarshal([]byte(v[:]), &as); e != nil {
			return as, e
		}
		return as, nil
	} else {
		return as, errors.New("返回数据类型有误，json无法解析")
	}
}