package user

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Id uint
}
