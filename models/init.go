package models

import (
	"server/db"
	"server/models/article"
	"server/models/friend"
	"server/models/user"
)

func Init() {
	tab := []interface{}{
		&user.User{},
		&article.Tag{},
		&article.Class{},
		&article.Article{},
		&friend.Friend{},
	}
	db.CreateTable(tab)
}
