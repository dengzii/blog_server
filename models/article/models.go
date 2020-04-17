package article

import (
	"github.com/dengzii/blog_server/models/base"
	"github.com/dengzii/blog_server/tools"
	"github.com/jinzhu/gorm"
)

type Article struct {
	base.CommonModel
	Title       string   `json:"title"`
	AuthorId    int      `json:"author_id"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Tag         Tag      `json:"tag" gorm:"ForeignKey:TagId"`
	TagId       int      `json:"tag_id"`
	Class       Category `json:"class" gorm:"ForeignKey:ClassId"`
	ClassId     int      `json:"class_id"`
	Likes       int      `json:"likes"`
	Comments    int      `json:"comments"`
	Views       int      `json:"views"`
	Display     bool     `json:"-" gorm:"default:true"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	tools.Log("Ready to create article,", article)
	//scope.SetColumn("ID", time.Now())
	return nil
}

func (article *Article) AfterCreate(scope *gorm.Scope) error {
	//tools.Log("Ready to create article,", article)
	//scope.SetColumn("ID", time.Now())
	return nil
}

func (article *Article) BeforeDelete(scope *gorm.Scope) error {
	//tools.Log("Ready to create article,", article)
	//scope.SetColumn("ID", time.Now())
	return nil
}

type Category struct {
	base.CommonModel
	Name         string `json:"name"`
	ArticleCount int    `json:"article_count"`
	Display      bool   `json:"display"`
}

type Tag struct {
	base.CommonModel
	ClassId      uint   `json:"class_id"`
	Name         string `json:"name"`
	ArticleCount int    `json:"article_count"`
	Display      bool   `json:"display"`
	Style        int    `json:"style"`
}
