package article

import "github.com/jinzhu/gorm"

type Class struct {
	gorm.Model
	Name         string
	ArticleCount int
	Display      bool
	ArticleID    uint
}
