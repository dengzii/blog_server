package article

import (
	"github.com/kataras/iris/context"
	"server/controllers"
	"server/models/article"
)

type tagJson struct {
	name  string
	style int
}

func AddTagController(ctx context.Context) (err error) {
	tagJson := &tagJson{}
	err = ctx.ReadJSON(tagJson)
	tag := article.AddTag(tagJson.name, tagJson.style)
	if tag == nil {
		err = controllers.NewControllerError("add tag failure.", 0)
	} else {
		_, err = ctx.JSON(controllers.SuccessResponse("add tag success", tag))
	}
	return err
}
