package article

import (
	"errors"
	"github.com/kataras/iris/context"
	"server/controllers"
	"server/models/article"
)

func GetArticle(ctx context.Context) (err error) {

	articles := article.GetArticle(0)
	responseJson := controllers.ErrorResponse(200, "success", articles)
	_, err = ctx.JSON(responseJson)
	return err
}

func AddArticle(ctx context.Context) (err error) {

	articleJson := &article.ArticleJson{}
	err = ctx.ReadJSON(articleJson)
	if err != nil {
		return err
	}
	art := article.AddArticle(articleJson)
	if art == nil {
		return errors.New("create article failure")
	}
	_, err = ctx.JSON(controllers.SuccessResponse("success", art))
	return err
}
