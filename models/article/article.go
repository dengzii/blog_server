package article

import (
	"github.com/dengzii/blog_server/db"
	"time"
)

func (newArticle *Article) Insert() (article *Article, err error) {
	desc := newArticle.Content
	if len(newArticle.Content) > 50 {
		desc = newArticle.Content[:50]
	}
	newArticle.Description = desc
	newArticle.CreatedAt = time.Now().Unix()
	db.Insert(newArticle)
	return newArticle, nil
}

func Insert(newArticle *Article) (article *Article, err error) {
	desc := newArticle.Content
	if len(newArticle.Content) > 50 {
		desc = newArticle.Content[:50]
	}
	newArticle.Description = desc
	newArticle.CreatedAt = time.Now().Unix()
	db.Insert(newArticle)
	return newArticle, nil
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
