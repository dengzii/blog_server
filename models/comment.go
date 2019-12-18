package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Avatar    string
	Name      string
	Email     string
	Site      string
	Content   string
	Reply     string
	Child     []Comment
	CommentID uint
	Display   bool
	ArticleID uint
}

type CommentJson struct {
	Id        uint
	Avatar    string
	Name      string
	Email     string
	Site      string
	Content   string
	Reply     string
	CommentID uint
	Display   bool
	ArticleID uint
}
