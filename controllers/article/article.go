package article

import (
	"errors"
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris/context"
	"time"
)

func GetArticles(ctx context.Context) (err error) {

	last, err := ctx.URLParamInt64("last")
	var category = ctx.URLParam("category")
	if last == -1 || err != nil {
		last = time.Now().Unix()
	}
	articles := article.GetArticles(last, category, 3)
	responseJson := controllers.SuccessResponse(articles)
	_, err = ctx.JSON(responseJson)
	return err
}

func GetArchive(ctx context.Context) (err error) {

	archive := article.GetArchive()
	response := controllers.SuccessResponse(archive)
	_, err = ctx.JSON(response)
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
	art, _ := article.AddArticle(newArticle)
	if art == nil {
		return errors.New("create article failure")
	}
	_, err = ctx.JSON(controllers.SuccessResponse(art))
	return err
}
