package article

import (
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris/context"
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
		_, err = ctx.JSON(controllers.SuccessResponse(tag))
	}
	return err
}

func GetTagsController(ctx context.Context) (err error) {

	tags := article.GetTags()
	_, err = ctx.JSON(controllers.SuccessResponse(tags))
	return
}
