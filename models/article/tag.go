package article

import (
	"github.com/jinzhu/gorm"
	"server/db"
)

type Tag struct {
	gorm.Model
	Belong       Class
	Name         string
	ArticleCount int
	Display      bool
}

func AddTag() (err error) {

	db.Insert(&Tag{})
	return err
}

func GetTag() (tag *Tag, err error) {

	return nil, nil
}
