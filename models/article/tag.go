package article

import (
	"github.com/dengzii/blog_server/db"
	"github.com/dengzii/blog_server/models/base"
)

type Tag struct {
	base.CommonModel
	ClassId      uint   `json:"class_id"`
	Name         string `json:"name"`
	ArticleCount int    `json:"article_count"`
	Display      bool   `json:"display"`
	Style        int    `json:"style"`
}

func AddTag(name string, style int) *Tag {

	tag := &Tag{
		ClassId:      0,
		Name:         name,
		ArticleCount: 0,
		Display:      true,
		Style:        style,
	}
	db.Insert(tag)
	return tag
}

func GetTags() (tag *[]Tag) {

	var tags []Tag
	db.Mysql.Find(&tags).Limit(5)
	return &tags
}
