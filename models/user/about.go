package user

import (
	"github.com/dengzii/blog_server/db"
	"github.com/dengzii/blog_server/models/base"
	"time"
)

type About struct {
	base.CommonModel
	Content string `json:"content"`
	Enable  bool   `json:"enable"`
}

func GetAbout() (about *About) {

	var abouts []*About
	db.Mysql.
		Where("enable = ?", 1).
		Order("updated_at desc").
		Limit(1).
		Find(&abouts)
	about = abouts[0]
	return
}

func AddAbout(about *About) {

	about.UpdatedAt = time.Now().Unix()
	db.Insert(&about)
	return
}
