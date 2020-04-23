package article

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dengzii/blog_server/db"
	"github.com/jinzhu/gorm"
	"time"
)

func (that *Article) Insert() (article *Article, err error) {
	desc := that.Content
	if len(that.Content) > 50 {
		desc = that.Content[:50]
	}
	that.Description = desc
	that.CreatedAt = time.Now().Unix()
	db.Insert(that)
	return that, nil
}

func Insert(newArticle *Article) (article *Article, err error) {
	desc := newArticle.Content
	if len(newArticle.Content) > 50 {
		desc = newArticle.Content[:50]
	}
	newArticle.Description = desc

	h := md5.New()
	h.Write([]byte(newArticle.Title))
	newArticle.ID = hex.EncodeToString(h.Sum(nil))

	//newArticle.CreatedAt = time.Now().Unix()
	//newArticle.UpdatedAt = newArticle.CreatedAt
	db.Insert(newArticle)
	return newArticle, nil
}

func GetArticles(from int64, category string, count int) (articles []*ArticleBase) {
	var article []*Article

	var query *gorm.DB
	query = db.Mysql.Order("updated_at desc")

	if len(category) > 0 && category != "All" {
		query = query.Where("category_name = ?", category)
	}

	query.
		Where("updated_at < ?", from).
		Limit(count).
		Find(&article)

	articles = make([]*ArticleBase, len(article))
	for i, v := range article {
		articles[i] = v.toArticleBase()
	}
	return articles
}

func GetArticle(id int) *ArticleBase {
	var article ArticleBase
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
