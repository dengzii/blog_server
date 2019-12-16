package article

import (
	"github.com/kataras/iris/context"
	"server/controllers"
	"server/models/article"
)

func GetArticle(ctx context.Context) (err error) {

	articles := article.GetArticle(10)
	responseJson := controllers.CommonJson(200, "success", articles)
	_, err = ctx.JSON(responseJson)
	return err
}

func AddArticle(ctx context.Context) (err error) {

	articleJson := &article.ArticleJson{}
	err = ctx.ReadJSON(articleJson)
	if err != nil {
		return
	}
	article.AddArticle(articleJson)
	return
}
