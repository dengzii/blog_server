package article

import (
	"github.com/dengzii/blog_server/models/base"
	"github.com/dengzii/blog_server/tools"
	"github.com/jinzhu/gorm"
	"time"
)

type ArticleBase struct {
	base.CommonModel

	Title        string `json:"title"`
	AuthorId     int    `json:"-"`
	AuthorName   string `json:"author_name"`
	Description  string `json:"description"`
	TagName      string `json:"tag_name"`
	CategoryName string `json:"category_name"`
	Likes        int    `json:"likes"`
	Views        int    `json:"views"`
}

type Article struct {
	ArticleBase

	Comments int      `json:"comments"`
	Tag      Tag      `json:"tag" gorm:"ForeignKey:TagId"`
	Category Category `json:"category" gorm:"ForeignKey:CategoryId"`
	Content  string   `json:"content"`
	Display  bool     `json:"-" gorm:"default:true"`
}

func (that *Article) BeforeCreate(scope *gorm.Scope) error {
	tools.Log("Ready to create article,", that)
	//scope.SetColumn("ID", time.Now())
	that.CreatedAt = time.Now().Unix()
	that.UpdatedAt = that.CreatedAt
	return nil
}

func (that *Article) AfterCreate(scope *gorm.Scope) error {

	//tools.Log("Ready to create article,", article)
	//scope.SetColumn("ID", time.Now())
	that.CreatedAt = time.Now().Unix()
	that.UpdatedAt = that.CreatedAt
	return nil
}

func (that *Article) BeforeDelete(scope *gorm.Scope) error {
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

func (that *Article) toArticleBase() (articleBase *ArticleBase) {
	articleBase = &ArticleBase{
		Title:        that.Title,
		AuthorName:   that.AuthorName,
		Description:  that.Description,
		TagName:      that.TagName,
		CategoryName: that.CategoryName,
		Likes:        that.Likes,
		Views:        that.Views,
	}
	articleBase.CreatedAt = that.CreatedAt
	articleBase.UpdatedAt = that.UpdatedAt
	return
}
