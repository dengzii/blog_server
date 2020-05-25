package user

import (
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/user"
	"github.com/kataras/iris/context"
)

func GetAbout(ctx context.Context) (err error) {

	_, err = ctx.JSON(controllers.SuccessResponse(user.GetAbout()))
	return nil
}

func AddAbout(ctx context.Context) (err error) {

	var about user.About
	err = ctx.ReadJSON(&about)
	user.AddAbout(&about)
	_, err = ctx.JSON(controllers.SuccessResponse(nil))
	return nil
}
