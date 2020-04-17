package article

import (
	"github.com/dengzii/blog_server/db"
)

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
