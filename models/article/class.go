package article

import "github.com/jinzhu/gorm"

type Class struct {
	gorm.Model
	Id           int `gorm:"AUTO_INCREMENT"`
	Name         string
	ArticleCount int
	Display      bool
}
