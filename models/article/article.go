package article

import (
	"github.com/jinzhu/gorm"
	"server/db"
	"time"
)

type Article struct {
	gorm.Model
	Id          uint `gorm:"AUTO_INCREMENT"`
	Title       string
	AuthorId    int
	Content     string
	Description string
	Tag         Tag
	Class       Class
	CreateDate  time.Time
	LastModify  time.Time
	Likes       int
	Comments    int
	Views       int
	Display     bool
}

type ArticleJson struct {
	Title      string
	AuthorId   int
	Content    string
	Tag        Tag
	Class      Class
	CreateDate time.Time
	LastModify time.Time
}

func AddArticle(json ArticleJson) {
	a := Article{
		Title:       json.Title,
		AuthorId:    json.AuthorId,
		Content:     json.Content,
		Description: json.Content[0:50],
		Tag:         json.Tag,
		Class:       json.Class,
		CreateDate:  time.Now(),
		LastModify:  time.Now(),
		Likes:       0,
		Comments:    0,
		Views:       0,
		Display:     true,
	}
	db.Insert(&a)
}

func GetArticle(id int) (article *Article) {
	db.Mysql.Where("id = ?", id).Find(article)
	return article
}

func DelArticle(id int) {

}

func UpdateArticle() {

}

func LikeArticle() {

}

func CommentArticle() {

}
