package controllers

import (
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris/context"
)

func HomeController(ctx context.Context) (err error) {

	response := ErrorResponse(400, "wrong password or username", nil)
	articles := article.GetArticleLatest(10)
	if articles != nil {
		response.Status = 200
		response.Msg = "success"
		response.Data = articles
	}
	_, err = ctx.JSON(response)
	return err
}
