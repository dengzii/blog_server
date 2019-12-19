package models

import (
	"server/db"
	"server/models/article"
	"server/models/friend"
	"server/models/user"
)

func Init() {
	tableArticle := &article.Article{}
	tableUser := &user.User{}
	tableTag := &article.Tag{}
	tableClass := &article.Category{}
	tableFriend := &friend.Friend{}

	tab := []interface{}{
		tableUser, tableTag, tableClass, tableFriend, tableArticle,
		&Comment{},
	}

	db.CreateTable(tab)
}
