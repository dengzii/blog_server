package article

import (
	"github.com/dengzii/blog_server/db"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name         string
	ArticleCount int
	Display      bool
}

func AddCategory(name string) (category *Category) {

	category = &Category{
		Name:         name,
		ArticleCount: 0,
		Display:      true,
	}
	db.Insert(category)
	return
}

func GetCategories() *[]Category {
	var categories []Category
	db.Mysql.Find(&categories).Limit(5)
	return &categories
}
