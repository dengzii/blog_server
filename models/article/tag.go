package article

import (
	"github.com/jinzhu/gorm"
	"server/db"
)

type Tag struct {
	gorm.Model
	ClassId      uint
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

func GetTag() (tag *Tag, err error) {
	return nil, nil
}
