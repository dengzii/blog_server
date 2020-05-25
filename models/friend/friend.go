package friend

import (
	"github.com/dengzii/blog_server/db"
	"github.com/jinzhu/gorm"
)

type Friend struct {
	gorm.Model
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Url     string `json:"url"`
	Email   string `json:"emil"`
	Avatar  string `json:"avatar"`
	Alt     string `json:"alt"`
	Display bool   `json:"-"`
}

func AddFriend(f *Friend) *Friend {
	f.Display = false
	db.Insert(f)
	return f
}

func GetFriend() (f []*Friend) {
	db.Mysql.Where("display = ?", 1).Find(&f)
	return
}
