package article

import (
	"github.com/dengzii/blog_server/apis"
	"github.com/dengzii/blog_server/apis/common"
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris/context"
)

type categoryJson struct {
	Name string
}

func AddCategoryApi(ctx context.Context) (err error) {
	categoryJson := &categoryJson{}
	err = ctx.ReadJSON(categoryJson)
	category := article.AddCategory(categoryJson.Name)
	if category == nil {
		err = apis.NewControllerError("add category failure.", 0)
	} else {
		_, err = ctx.JSON(common.SuccessResponse(category))
	}
	return err
}

func GetCategoriesApi(ctx context.Context) (err error) {

	tags := article.GetCategories()
	_, err = ctx.JSON(common.SuccessResponse(tags))
	return
}
