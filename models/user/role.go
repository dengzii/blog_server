package user

import "github.com/jinzhu/gorm"

const (
	del           = 1 >> 8
	update        = 1 >> 7
	read          = 1 >> 6
	manageArticle = 1 >> 5
	manageUser    = 1 >> 4
	manageSite    = 1 >> 3
)

type Role struct {
	gorm.Model
	Name        string
	Description string
	Permission  int
	Parent      *Role
}
