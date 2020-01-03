package soup_service

import (
	"fmt"
	"github.com/lawtech0902/go_gin_blog/backend/models"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/utils"
	"github.com/unknwon/com"
	"time"
)

type Soup struct {
	ID      int    `json:"id" db:"id"`
	Content string `json:"content" db:"content" binding:"required"`
}

type Soups struct {
	Items []Soup `json:"items"`
	Total int    `json:"total"`
}

var (
	db = models.DB
)

func (s *Soup) Create() (Soup, error) {
	r, err := db.Exec("insert into blog_soup (content) values (?)", s.Content)
	if err != nil {
		return Soup{}, err
	}
	
	id, _ := r.LastInsertId()
	return Soup{ID: int(id), Content: s.Content}, nil
}

func (s *Soup) Delete() error {
	_, err := db.Exec("delete from blog_soup where id=?", s.ID)
	return err
}

func (s *Soup) Edit() error {
	_, err := db.Exec("update blog_soup set content=? where id=?", s.Content, s.ID)
	return err
}

func (s *Soup) GetOne() (Soup, error) {
	var soup Soup
	err := db.Get(&soup, "select * from blog_soup where id=?", s.ID)
	return soup, err
}

func (s Soup) GetRandOne() (Soup, error) {
	var soup Soup
	querySql := `select t1.* from blog_soup as t1 join (select round(rand() * ((select max(id) FROM blog_soup) - (select min(id) from blog_soup)) + (select min(id) from blog_soup)) as id
) as t2 where t1.id >= t2.id limit 1`
	err := db.Get(&soup, querySql)
	return soup, err
}

func (s Soup) GetAll(limit, page string) (data Soups, err error) {
	baseSql := "select %s from blog_soup s"
	var l, p, total int
	if limit != "" && page != "" {
		l = com.StrTo(limit).MustInt()
		p = com.StrTo(page).MustInt()
	}
	key := soupCacheKey(l, p)
	
	cacheData, err := getSoupCache(key)
	if err != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 读取缓存失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
	}
	if cacheData.Total != 0 {
		return cacheData, nil
	}
	
	soups := make([]Soup, 0)
	offset := (p - 1) * l
	selectSql := fmt.Sprintf(baseSql, "s.*") + fmt.Sprintf(" order by s.id desc limit %d offset %d", l, offset)
	if err = db.Select(&soups, selectSql); err != nil {
		return
	}
	
	if err = db.Get(&total, fmt.Sprintf(baseSql, "count(1)")); err != nil {
		return
	}
	
	data.Total = total
	data.Items = soups
	
	if err = setSoupCache(key, data); err != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 写入缓存失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
	}
	
	return
}
