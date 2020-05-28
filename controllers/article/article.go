package article

import (
	"errors"
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"time"
)

func GetArticles(ctx context.Context) (err error) {

	last, err := ctx.URLParamInt64("last")
	var category = ctx.URLParam("category")
	if last == -1 || err != nil {
		last = time.Now().Unix()
	}
	articles := article.GetArticles(last, category, 10)
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
	if err != nil || id <= 0 {
		err = errors.New("bad request")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	articles := article.GetArticle(id)
	if articles == nil {
		err = errors.New("article not found")
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

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

func ViewArticle(ctx context.Context) (err error) {

	articleId, err := ctx.URLParamInt("id")
	if err != nil || articleId <= 0 {
		err = errors.New("bad request")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = article.ViewArticle(articleId)
	return err
}

func LikeArticle(ctx context.Context) (err error) {

	return err
}

func CommentArticle(ctx context.Context) (err error) {

	return err
}
