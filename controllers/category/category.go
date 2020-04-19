package category

import (
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris/context"
)

type categoryJson struct {
	Name string
}

func AddCategoryController(ctx context.Context) (err error) {
	categoryJson := &categoryJson{}
	err = ctx.ReadJSON(categoryJson)
	category := article.AddCategory(categoryJson.Name)
	if category == nil {
		err = controllers.NewControllerError("add category failure.", 0)
	} else {
		_, err = ctx.JSON(controllers.SuccessResponse(category))
	}
	return err
}

func GetCategoriesController(ctx context.Context) (err error) {

	tags := article.GetCategories()
	_, err = ctx.JSON(controllers.SuccessResponse(tags))
	return
}
