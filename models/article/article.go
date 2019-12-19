package article

import (
	"github.com/jinzhu/gorm"
	"server/db"
)

type Article struct {
	gorm.Model
	Title       string
	AuthorId    int
	Content     string
	Description string
	Tag         Tag `gorm:"ForeignKey:TagId"`
	TagId       int
	Class       Category `gorm:"ForeignKey:ClassId"`
	ClassId     int
	Likes       int
	Comments    int
	Views       int
	Display     bool `gorm:"default:true"`
}

type ArticleJson struct {
	Title    string
	AuthorId int
	Content  string
	TagId    int
	ClassId  int
}

func AddArticle(json *ArticleJson) (article *Article) {

	desc := json.Content
	if len(json.Content) > 50 {
		desc = json.Content[:50]
	}
	article = &Article{
		Title:       json.Title,
		AuthorId:    json.AuthorId,
		Content:     json.Content,
		Description: desc,
		TagId:       json.TagId,
		ClassId:     json.ClassId,
		Likes:       0,
		Comments:    0,
		Views:       0,
		Display:     true,
	}
	db.Insert(article)
	return article
}

func GetArticleLatest(count int) (article []*Article) {

	db.Mysql.Find(&article).Order("create_at", true).Limit(count)
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
