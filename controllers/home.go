package controllers

import (
	"github.com/kataras/iris/context"
	"server/models/article"
)

func HomeController(ctx context.Context) (err error) {

	response := ErrorResponse(400, "wrong password or username", nil)
	articles := article.GetArticleLatest(20)
	if articles != nil {
		response.Data = articles
	}
	_, err = ctx.JSON(response)
	return err
}
