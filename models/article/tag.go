package article

import (
	"github.com/dengzii/blog_server/db"
)

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
