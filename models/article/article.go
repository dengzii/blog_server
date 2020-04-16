package article

import (
	"github.com/dengzii/blog_server/db"
	"github.com/dengzii/blog_server/models/base"
	"time"
)

type Article struct {
	base.CommonModel
	Title       string   `json:"title"`
	AuthorId    int      `json:"author_id"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Tag         Tag      `json:"tag" gorm:"ForeignKey:TagId"`
	TagId       int      `json:"tag_id"`
	Class       Category `json:"class" gorm:"ForeignKey:ClassId"`
	ClassId     int      `json:"class_id"`
	Likes       int      `json:"likes"`
	Comments    int      `json:"comments"`
	Views       int      `json:"views"`
	Display     bool     `json:"-" gorm:"default:true"`
}

type Json struct {
	Title    string
	AuthorId int
	Content  string
	TagId    int
	ClassId  int
}

func AddArticle(new *Article) (article *Article) {
	desc := new.Content
	if len(new.Content) > 50 {
		desc = new.Content[:50]
	}
	new.Description = desc
	new.CreatedAt = time.Now().Unix()
	db.Insert(new)
	return new
}

func GetArticleLatest(count int) (article []*Article) {
	db.Mysql.Find(&article).Order("created_at", true).Limit(count)
	return article
}

func GetArticle(id int) *Article {
	var article Article
	db.Mysql.Where("id = ?", id).Find(&article).Limit(1)
	return &article
}

func DelArticle(id int) {

}

func UpdateArticle() {

}

func LikeArticle() {

}

func CommentArticle() {

}
