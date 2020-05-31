package article

import (
	"errors"
	"github.com/dengzii/blog_server/apis/common"
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"time"
)

func GetArticlesApi(ctx context.Context) (err error) {

	last, err := ctx.URLParamInt64("last")
	var category = ctx.URLParam("category")
	if last == -1 || err != nil {
		last = time.Now().Unix()
	}
	articles := article.GetArticles(last, category, 10)
	responseJson := common.SuccessResponse(articles)
	_, err = ctx.JSON(responseJson)
	return err
}

func GetArchiveApi(ctx context.Context) (err error) {

	archive := article.GetArchive()
	response := common.SuccessResponse(archive)
	_, err = ctx.JSON(response)
	return err
}

func GetArticleApi(ctx context.Context) (err error) {

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

	responseJson := common.SuccessResponse(articles)
	_, err = ctx.JSON(responseJson)
	return
}

func AddArticleApi(ctx context.Context) (err error) {

	newArticle := &article.Article{}
	err = ctx.ReadJSON(newArticle)
	if err != nil {
		return err
	}
	art, _ := article.AddArticle(newArticle)
	if art == nil {
		return errors.New("create article failure")
	}
	_, err = ctx.JSON(common.SuccessResponse(art))
	return err
}

func ViewArticleApi(ctx context.Context) (err error) {

	articleId, err := ctx.URLParamInt("id")
	if err != nil || articleId <= 0 {
		err = errors.New("bad request")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = article.ViewArticle(articleId)
	return err
}

func LikeArticleApi(ctx context.Context) (err error) {

	return err
}

func CommentArticleApi(ctx context.Context) (err error) {

	return err
}
