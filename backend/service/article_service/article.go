package article_service

import (
	"fmt"
	"github.com/lawtech0902/go_gin_blog/backend/models"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/gredis"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/utils"
	"github.com/lawtech0902/go_gin_blog/backend/service/category_service"
	"github.com/lawtech0902/go_gin_blog/backend/service/comment_service"
	"github.com/lawtech0902/go_gin_blog/backend/service/tag_service"
	"time"
)

type Article struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required,max=32"`
	Content     string `json:"content" db:"content" binding:"required"`
	Html        string `json:"html" db:"html" binding:"required"`
	CategoryID  int    `json:"category_id" db:"category_id" binding:"required"`
	TagID       []int  `json:"tag_id" binding:"required"`
	CreatedTime string `json:"created_time" db:"created_time"`
	UpdatedTime string `json:"updated_time" db:"updated_time"`
	Status      string `json:"status" db:"status" binding:"required"`
}

type Articles struct {
	Items []Article `json:"items"`
	Total int       `json:"total"`
}

type ArticleDetail struct {
	A     Article                   `json:"article"`
	C     category_service.Category `json:"category"`
	Tags  []tag_service.Tag         `json:"tags"`
	Views uint8                     `json:"views"`
}

var (
	db = models.DB
	es = models.ESConn
)

// article crud 操作实现
func (a *Article) Create() (Article, error) {
	createdTime := time.Now().Format(setting.AppInfo.TimeFormat)
	
	r, err := db.Exec("insert into blog_article (title, content, html, category_id, created_time, status) values (?, ?, ?, ?, ?, ?)",
		a.Title,
		a.Content,
		a.Html,
		a.CategoryID,
		createdTime,
		a.Status)
	if err != nil {
		return Article{}, err
	}
	
	articleID, _ := r.LastInsertId()
	if len(a.TagID) > 0 {
		for _, tagID := range a.TagID {
			_, err := db.Exec("insert into blog_tag_article(tag_id, article_id) values (?, ?)", tagID, articleID)
			if err != nil {
				return Article{}, err
			}
		}
	}
	
	article := Article{
		ID:          int(articleID),
		Title:       a.Title,
		Content:     a.Content,
		Html:        a.Html,
		CategoryID:  a.CategoryID,
		TagID:       a.TagID,
		CreatedTime: createdTime,
		UpdatedTime: a.UpdatedTime,
		Status:      a.Status,
	}
	
	if article.Status == "published" {
		if err := article.IndexBlog(); err != nil {
			utils.WriteErrorLog(fmt.Sprintf("[ %s ] 存入elastic出错, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
		}
	}
	
	return article, nil
}

func (a *Article) Edit() error {
	updatedTime := time.Now().Format(setting.AppInfo.TimeFormat)
	
	if _, err := db.Exec("update blog_article set title=?, content=?, html=?, category_id=?, updated_time=?, status=? where id=?",
		a.Title,
		a.Content,
		a.Html,
		a.CategoryID,
		updatedTime,
		a.Status,
		a.ID); err != nil {
		return err
	}
	
	if _, err := db.Exec("delete from blog_tag_article where article_id=?", a.ID); err != nil {
		return err
	}
	
	if len(a.TagID) > 0 {
		for _, tagID := range a.TagID {
			_, err := db.Exec("insert into blog_tag_article(tag_id, article_id) values (?, ?)", tagID, a.ID)
			if err != nil {
				return err
			}
		}
	}
	
	if a.Status == "published" {
		if err := a.IndexBlog(); err != nil {
			utils.WriteErrorLog(fmt.Sprintf("[ %s ] 从elastic更新出错, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
		}
	} else {
		if err := a.DeleteFromES(); err != nil {
			utils.WriteErrorLog(fmt.Sprintf("[ %s ] 从elastic删除出错, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
		}
	}
	
	return nil
}

func (a *Article) Delete() error {
	if _, err := db.Exec("delete from blog_tag_article where article_id=?", a.ID); err != nil {
		return err
	}
	
	if _, err := db.Exec("delete from blog_article where id=?", a.ID); err != nil {
		return err
	}
	
	viewKey := a.ViewKey()
	if err := gredis.DelKey(viewKey); err != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 删除阅读量失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
	}
	// 从ES中删除
	if err := a.DeleteFromES(); err != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 从elastic中删除出错, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
	}
	
	return nil
}

func (a Article) GetOne(opts ...Option) (ArticleDetail, error) {
	options := newOptions(opts...)
	
	var one Article
	
	if err := db.Get(&one, "select * from blog_article where id=?", a.ID); err != nil {
		return ArticleDetail{}, err
	}
	
	category, _ := GetCategoryByID(one.CategoryID)
	tags, _ := GetTagsByArticleID(a.ID)
	
	viewKey := one.ViewKey()
	n, err := getViews(viewKey)
	if err != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 获取阅读量失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
	}
	
	if !options.Admin {
		if err := addView(viewKey); err != nil {
			utils.WriteErrorLog(fmt.Sprintf("[ %s ] 添加阅读量失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
		}
	}
	
	return ArticleDetail{
		A:     one,
		C:     category,
		Tags:  tags,
		Views: n,
	}, nil
}

func (a Article) GetAll(opts ...Option) (data Articles, err error) {
	baseSql := "select %s from blog_article a"
	data, err = genArticles(baseSql, opts...)
	return
}

func (a Article) GetCommentsByArticle() (comment_service.ArticleComments, error) {
	var rootComments []comment_service.RootComment
	roots, err := a.GetRootCommentsByArticle()
	if err != nil {
		return comment_service.ArticleComments{}, err
	}
	
	for _, value := range roots {
		comments, err := value.GetChildren()
		if err != nil {
			return comment_service.ArticleComments{}, err
		}
		
		rootComments = append(rootComments, comment_service.RootComment{Comment: value, Children: comments})
	}
	
	count, _ := a.getCommentsCount()
	return comment_service.ArticleComments{Items: rootComments, Total: count}, nil
}

func (a Article) GetRootCommentsByArticle() ([]comment_service.Comment, error) {
	comments := make([]comment_service.Comment, 0)
	err := db.Select(&comments, "select * from blog_comment where article_id=? and root_id is null", a.ID)
	return comments, err
}

func (a Article) getCommentsCount() (uint, error) {
	var count uint
	err := db.Get(&count, "select count(id) from blog_comment where article_id=?", a.ID)
	return count, err
}

func GetArticlesByCategory(opts ...Option) (data Articles, err error) {
	options := newOptions(opts...)
	baseSql := "SELECT %s FROM blog_article a INNER JOIN blog_category c ON a.category_id=c.id AND c.category_name=" + "'" + options.C + "'" + ""
	data, err = genArticles(baseSql, opts...)
	return
}

func GetArticlesByTag(opts ...Option) (data Articles, err error) {
	options := newOptions(opts...)
	baseSql := "SELECT %s FROM blog_article a  INNER JOIN blog_tag_article ta ON a.id=ta.article_id INNER JOIN blog_tag t ON ta.tag_id=t.id AND t.tag_name=" + "'" + options.T + "'" + ""
	data, err = genArticles(baseSql, opts...)
	return
}

func SearchArticle(key, status string, opts ...Option) (data Articles, err error) {
	var baseSql string
	if status == "" {
		baseSql = `SELECT %s FROM blog_article a WHERE a.title LIKE '%%` + key + `%%'`
	} else {
		baseSql = `SELECT %s FROM blog_article a WHERE a.title LIKE '%%` + key + `%%' AND a.status='` + status + `'`
	}
	
	data, err = genArticles(baseSql, opts...)
	return
}

func GetTagsByArticleID(articleID int) ([]tag_service.Tag, error) {
	var tags []tag_service.Tag
	if err := db.Select(&tags, "select t.* from blog_tag t right join blog_tag_article ta on t.id=ta.tag_id where ta.article_id=?", articleID); err != nil {
		return nil, err
	}
	
	return tags, nil
}

func GetCategoryByID(id int) (category_service.Category, error) {
	var c category_service.Category
	if err := db.Get(&c, "select * from blog_category where id=?", id); err != nil {
		return category_service.Category{}, err
	}
	
	return c, nil
}

func genArticles(baseSql string, opts ...Option) (data Articles, err error) {
	options := newOptions(opts...)
	key := articleCacheKey(options)
	if !options.Search {
		cacheData, err := getArticleCache(key)
		if err != nil {
			utils.WriteErrorLog(fmt.Sprintf("[ %s ] 读取缓存失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
		}
		if cacheData.Total != 0 {
			return cacheData, nil
		}
	}
	
	articles := make([]Article, 0)
	
	var f string
	if !options.Admin {
		f = " WHERE a.`status`='published'"
	}
	offset := (options.Page - 1) * options.Limit
	selectSql := fmt.Sprintf(baseSql, "a.id, a.title, a.created_time, a.updated_time, a.`status`") + f + fmt.Sprintf(" ORDER BY a.id DESC limit %d offset %d", options.Limit, offset)
	
	if err = db.Select(&articles, selectSql); err != nil {
		return
	}
	
	var total int
	if err = db.Get(&total, fmt.Sprintf(baseSql, "count(1)")+f); err != nil {
		return
	}
	
	data.Total = total
	data.Items = articles
	
	if !options.Search {
		if err := setArticleCache(key, data); err != nil {
			utils.WriteErrorLog(fmt.Sprintf("[ %s ] 写入缓存失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
		}
	}
	
	return
}
