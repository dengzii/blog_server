package category

import (
	"github.com/kataras/iris/context"
	"server/controllers"
	"server/models/article"
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
		_, err = ctx.JSON(controllers.SuccessResponse("add category success", category))
	}
	return err
}

func GetCategoriesController(ctx context.Context) (err error) {

	tags := article.GetCategories()
	_, err = ctx.JSON(controllers.SuccessResponse("get category success", tags))
	return
}
