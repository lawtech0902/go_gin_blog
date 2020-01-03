package category_service

import "github.com/lawtech0902/go_gin_blog/backend/models"

type Category struct {
	ID           int    `json:"id" db:"id"`
	CategoryName string `json:"category_name" db:"category_name" binding:"required,max=16"`
}

var (
	db = models.DB
)

// category crud 操作实现
func (c *Category) Create() (Category, error) {
	r, err := db.Exec("insert into blog_category (category_name) values (?)", c.CategoryName)
	if err != nil {
		return Category{}, err
	}
	
	id, _ := r.LastInsertId()
	return Category{
		ID:           int(id),
		CategoryName: c.CategoryName,
	}, err
}

func (c *Category) Delete() error {
	_, err := db.Exec("delete from blog_category where id=?", c.ID)
	return err
}

func (c *Category) Edit() error {
	_, err := db.Exec("update blog_category set category_name=? where id=?", c.CategoryName, c.ID)
	return err
}

func (c Category) GetOne() (Category, error) {
	var cg Category
	err := db.Get(&cg, "select * from blog_category where id=?", c.ID)
	return cg, err
}

func (c Category) GetAll() ([]Category, error) {
	cgs := make([]Category, 0)
	err := db.Select(&cgs, "select * from blog_category")
	return cgs, err
}

