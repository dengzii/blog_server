package article

import (
	"github.com/dengzii/blog_server/db"
	"github.com/dengzii/blog_server/models/base"
)

type Category struct {
	base.CommonModel
	Name         string `json:"name"`
	ArticleCount int    `json:"article_count"`
	Display      bool   `json:"display"`
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
