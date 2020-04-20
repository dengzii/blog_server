package article

import (
	"errors"
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris/context"
	"time"
)

func GetArticleLatest(ctx context.Context) (err error) {

	from, err := ctx.URLParamInt64("from")
	if from == -1 || err != nil {
		from = time.Now().Unix()
	}
	articles := article.GetArticleLatest(from, 3)
	responseJson := controllers.SuccessResponse(articles)
	_, err = ctx.JSON(responseJson)
	return err
}

func GetArticle(ctx context.Context) (err error) {

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		return
	}
	articles := article.GetArticle(id)
	responseJson := controllers.SuccessResponse(articles)
	_, err = ctx.JSON(responseJson)
	return
}

func AddArticle(ctx context.Context) (err error) {

	newArticle := &article.Article{}
	err = ctx.ReadJSON(newArticle)
	if err != nil {
		return err
	}
	art, _ := newArticle.Insert()
	if art == nil {
		return errors.New("create article failure")
	}
	_, err = ctx.JSON(controllers.SuccessResponse(art))
	return err
}
