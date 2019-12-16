package friend

import (
	"github.com/jinzhu/gorm"
	"server/db"
)

type Friend struct {
	gorm.Model
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
	Email  string `json:"emil"`
	Avatar string `json:"avtar"`
	Alt    string `json:"alt"`
}

func AddFriend(f Friend) *Friend {

	db.Insert(f)
	return &f
}

func GetFriend(f Friend) *Friend {
	return nil
}
