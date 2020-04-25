package models

import (
	"github.com/dengzii/blog_server/db"
	"github.com/dengzii/blog_server/models/article"
	"github.com/dengzii/blog_server/models/friend"
	"github.com/dengzii/blog_server/models/user"
)

func Init() {
	tableArticle := &article.Article{}
	tableUser := &user.User{}
	tableTag := &article.Tag{}
	tableClass := &article.Category{}
	tableFriend := &friend.Friend{}
	tableArchive := &article.Archive{}

	tab := []interface{}{
		tableUser, tableTag, tableClass, tableFriend, tableArticle, tableArchive,
		&Comment{}, &user.About{},
	}

	db.CreateTable(tab)
}
