package article

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/dengzii/blog_server/db"
	"github.com/jinzhu/gorm"
	"time"
)

func AddArticle(newArticle *Article) (article *Article, err error) {

	desc := newArticle.Content
	if len(newArticle.Content) > 50 {
		desc = newArticle.Content[:50]
	}
	newArticle.Description = desc

	h := md5.New()
	h.Write([]byte(newArticle.Title))
	newArticle.CID = hex.EncodeToString(h.Sum(nil))

	newArticle.CreatedAt = time.Now().Unix()
	newArticle.UpdatedAt = newArticle.CreatedAt

	if db.Insert(newArticle) {
		db.Insert(&Archive{
			CreatedAt: newArticle.CreatedAt,
			ID:        newArticle.ID,
			Title:     newArticle.Title,
		})
	}
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

func GetArticle(id int) *Article {
	var article Article
	db.Mysql.Where("id = ?", id).Attrs(nil).FirstOrInit(&article)
	if len(article.Title) == 0 {
		return nil
	}
	return &article
}

func ViewArticle(id int) (err error) {
	var article Article
	db.Mysql.Where("id = ?", id).First(&article)
	if len(article.Title) == 0 {
		err = errors.New("article not found")
		return
	}
	article.Views += 1
	db.Mysql.Model(&article).Update("views", article.Views)
	return
}

func GetArchive() (archive []*Archive) {
	db.Mysql.Order("created_at desc").Find(&archive)
	return
}

func DelArticle(id int) {

}

func UpdateArticle() {

}

func LikeArticle() {

}

func CommentArticle() {

}
