package tag_service

import "github.com/lawtech0902/go_gin_blog/backend/models"

type Tag struct {
	ID      int    `json:"id" db:"id"`
	TagName string `json:"tag_name" db:"tag_name" binding:"required,max=16"`
}

var (
	db = models.DB
)

// tag crud 操作实现
func (t *Tag) Create() (Tag, error) {
	r, err := db.Exec("insert into blog_tag (tag_name) values (?)", t.TagName)
	if err != nil {
		return Tag{}, err
	}
	id, _ := r.LastInsertId()
	return Tag{int(id), t.TagName}, nil
}

func (t *Tag) Delete() error {
	_, err := db.Exec("delete from blog_tag where id=?", t.ID)
	return err
}

func (t *Tag) Edit() error {
	_, err := db.Exec("update blog_tag set tag_name=? where id=?", t.TagName, t.ID)
	return err
}

func (t Tag) GetOne() (Tag, error) {
	var tag Tag
	err := db.Get(&tag, "select * from blog_tag where id=?", t.ID)
	return tag, err
}

func (t Tag) GetAll() ([]Tag, error) {
	tags := make([]Tag, 0)
	err := db.Select(&tags, "select * from blog_tag")
	return tags, err
}
