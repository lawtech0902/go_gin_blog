package soup_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/gredis"
)

func soupCacheKey(limit, page int) string {
	if limit == 0 || page == 0 {
		return fmt.Sprintf("soup_%d_%d", 10, 1)
	}
	
	return fmt.Sprintf("soup_%d_%d", limit, page)
}

func setSoupCache(key string, value Soups) error {
	marshal, _ := json.Marshal(value)
	err := gredis.SetKey(key, marshal, gredis.SetTimeout(true))
	return err
}

func getSoupCache(key string) (s Soups, err error) {
	data, err := gredis.GetKey(key)
	if err != nil || data == nil {
		return s, err
	}
	
	v, ok := data.([]uint8)
	if ok {
		if err := json.Unmarshal([]byte(v[:]), &s); err != nil {
			return s, err
		}
		return s, nil
	} else {
		return s, errors.New("返回数据类型有误，json无法解析")
	}
}
