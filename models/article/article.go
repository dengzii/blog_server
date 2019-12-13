package article

import (
	"server/models/user"
)

type Article struct {
	Id          int64
	Title       string
	Author      user.User
	Content     string
	Description string
	Tag         Tag
	Class       Class
	CreateDate  int64
	LastModify  int64
	Likes       int
	Comments    int
	Views       int
	Display     bool
}

func AddArticle() {

}

func GetArticle(id int) {

}

func DelArticle(id int) {

}

func UpdateArticle() {

}

func LikeArticle() {

}

func CommentArticle() {

}
