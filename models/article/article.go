package article

import (
	"github.com/jinzhu/gorm"
	"server/db"
	"time"
)

type Article struct {
	gorm.Model
	Title       string
	AuthorId    int
	Content     string
	Description string
	Tag         Tag `gorm:"ForeignKey:TagId"`
	TagId       int
	Class       Class `gorm:"ForeignKey:ClassId"`
	ClassId     int
	Likes       int
	Comments    int
	Views       int
	Display     bool `gorm:"default:'true'"`
}

type ArticleJson struct {
	Title      string
	AuthorId   int
	Content    string
	TagId      int
	ClassId    int
	CreateDate time.Time
	LastModify time.Time
}

func AddArticle(json *ArticleJson) {
	a := Article{
		Title:       json.Title,
		AuthorId:    json.AuthorId,
		Content:     json.Content,
		Description: json.Content[0:50],
		TagId:       json.TagId,
		ClassId:     json.ClassId,
		Likes:       0,
		Comments:    0,
		Views:       0,
		Display:     true,
	}
	db.Insert(&a)
}

func GetArticleLatest(count int) (article []*Article) {

	db.Mysql.Find(article).Order("create_at", true).Limit(count)
	return article
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
