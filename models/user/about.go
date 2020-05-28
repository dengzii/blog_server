package user

import (
	"github.com/dengzii/blog_server/db"
	"github.com/dengzii/blog_server/models/base"
	"time"
)

type About struct {
	base.CommonModel
	Content string `json:"content" gorm:"type:TEXT"`
	Enable  bool   `json:"-"`
}

func GetAbout() (about *About) {

	var abouts []*About
	db.Mysql.
		Where("enable = ?", 1).
		Order("updated_at desc").
		Limit(1).
		Find(&abouts)
	if len(abouts) == 0 {
		about = &About{}
		return
	}
	about = abouts[0]
	return
}

func AddAbout(about *About) (err error) {

	about.UpdatedAt = time.Now().Unix()
	db.Insert(&about)
	return
}
