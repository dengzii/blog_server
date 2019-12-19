package controllers

import (
	"github.com/kataras/iris/context"
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
		err = NewControllerError("add category failure.", 0)
	} else {
		_, err = ctx.JSON(SuccessResponse("add category success", category))
	}
	return err
}

func GetCategoriesController(ctx context.Context) (err error) {

	tags := article.GetCategories()
	_, err = ctx.JSON(SuccessResponse("get category success", tags))
	return
}
