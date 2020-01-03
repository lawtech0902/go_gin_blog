package comment_service

import (
	"database/sql"
	"fmt"
	"github.com/lawtech0902/go_gin_blog/backend/models"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/utils"
	"github.com/unknwon/com"
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

type Comment struct {
	ID          uint          `json:"id" db:"id"`
	UserName    string        `json:"username" db:"username" binding:"required,max=16"`
	IsAuthor    bool          `json:"is_author" db:"is_author"`
	ParentID    sql.NullInt64 `json:"parent_id" db:"parent_id"` // 回复某条评论的ID
	RootID      sql.NullInt64 `json:"root_id" db:"root_id"`     // 根评论ID
	ArticleID   uint          `json:"article_id" db:"article_id" binding:"required"`
	Content     string        `json:"content" db:"content" binding:"required"`
	CreatedTime string        `json:"created_time" db:"created_time"`
}

type RootComment struct {
	Comment
	Children []Comment `json:"children"`
}

type ArticleComments struct {
	Items []RootComment `json:"items"`
	Total uint          `json:"total"`
}

type CommentAndArticle struct {
	C Comment `json:"comment"`
	A Article `json:"article"`
}

type AllComments struct {
	Items []CommentAndArticle `json:"items"`
	Total uint                `json:"total"`
}

var db = models.DB

func (c *Comment) Create() (Comment, error) {
	var isAuthor int
	createTime := time.Now().Format(setting.AppInfo.TimeFormat)
	if c.IsAuthor {
		isAuthor = 1
	} else {
		isAuthor = 0
	}
	
	r, err := db.Exec("insert into blog_comment (username,is_author,parent_id,root_id,article_id,content, created_time) values (?,?,?,?,?,?,?)",
		c.UserName, isAuthor, c.ParentID, c.RootID, c.ArticleID, c.Content, createTime)
	if err != nil {
		return Comment{}, err
	}
	
	commentID, _ := r.LastInsertId()
	return Comment{
		ID:          uint(commentID),
		UserName:    c.UserName,
		IsAuthor:    c.IsAuthor,
		ParentID:    c.ParentID,
		RootID:      c.RootID,
		ArticleID:   c.ArticleID,
		Content:     c.Content,
		CreatedTime: createTime,
	}, nil
}

func (c *Comment) Delete() error {
	_, err := db.Exec("delete from blog_comment where id=?", c.ID)
	return err
}

func (c *Comment) GetOne() (Comment, error) {
	var comment Comment
	err := db.Get(&comment, "select * from blog_comment where id=?", c.ID)
	return comment, err
}

func (c Comment) GetAll(limit, page string) (data AllComments, err error) {
	var (
		l, p        int
		comments    []Comment
		allComments []CommentAndArticle
		total       uint
	)
	
	baseSql := "select %s from blog_comment s"
	
	if limit != "" && page != "" {
		p = com.StrTo(page).MustInt()
		l = com.StrTo(limit).MustInt()
	}
	
	key := commentCacheKey(l, p)
	
	cacheData, err := getCommentCache(key)
	if err != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 读取缓存失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
	}
	if cacheData.Total != 0 {
		return cacheData, nil
	}
	
	offset := (p - 1) * l
	selectSql := fmt.Sprintf(baseSql, "s.*") + fmt.Sprintf("order by s.id desc limit %d offset %d", l, offset)
	if err = db.Select(&comments, selectSql); err != nil {
		return
	}
	
	for _, value := range comments {
		article, _ := GetArticleByID(int(value.ArticleID))
		allComments = append(allComments, CommentAndArticle{value, article})
	}
	
	if err = db.Get(&total, fmt.Sprintf(baseSql, "count(id)")); err != nil {
		return
	}
	
	data.Total = total
	data.Items = allComments
	
	if err := setCommentCache(key, data); err != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 写入缓存失败, %v\n", time.Now().Format(setting.AppInfo.TimeFormat), err))
	}
	
	return
}

func (c *Comment) GetChildren() ([]Comment, error) {
	comments := make([]Comment, 0)
	err := db.Select(&comments, "select * from blog_comment where root_id=?", c.ID)
	return comments, err
}

func GetCommentsByArticleName(articleName string) (data AllComments, err error) {
	var (
		comments    []Comment
		allComments []CommentAndArticle
	)
	
	err = db.Select(&comments, "select c.* from blog_comment c inner join blog_article a on c.article_id=a.id where a.title=?", articleName)
	if err != nil {
		return
	}
	
	for _, value := range comments {
		allComments = append(allComments, CommentAndArticle{value, Article{ID: int(value.ArticleID), Title: articleName}})
	}
	
	data = AllComments{allComments, uint(len(comments))}
	return
}

func GetArticleByID(articleID int) (Article, error) {
	var article Article
	err := db.Get(&article, "select id, title from blog_article where id=?", articleID)
	return article, err
}
