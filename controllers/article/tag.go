package article

import (
	"github.com/kataras/iris/context"
	"server/controllers"
	"server/models/article"
)

type tagJson struct {
	Name  string `json:"name"`
	Style int    `json:"style"`
}

func AddTagController(ctx context.Context) (err error) {
	tagJson := &tagJson{}
	err = ctx.ReadJSON(tagJson)
	tag := article.AddTag(tagJson.Name, tagJson.Style)
	if tag == nil {
		err = controllers.NewControllerError("add tag failure.", 0)
	} else {
		_, err = ctx.JSON(controllers.SuccessResponse("add tag success", tag))
	}
	return err
}

func GetTagsController(ctx context.Context) (err error) {

	tags := article.GetTags()
	_, err = ctx.JSON(controllers.SuccessResponse("get tags success", tags))
	return
}
