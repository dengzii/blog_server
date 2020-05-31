package user

import (
	"github.com/dengzii/blog_server/apis/common"
	"github.com/dengzii/blog_server/models/user"
	"github.com/kataras/iris/context"
)

func GetAbout(ctx context.Context) (err error) {

	_, err = ctx.JSON(common.SuccessResponse(user.GetAbout()))
	return nil
}

func AddAbout(ctx context.Context) (err error) {

	about := &user.About{}

	err = ctx.ReadJSON(&about)
	println("=>" + about.Content)
	user.AddAbout(about)
	_, err = ctx.JSON(common.SuccessResponse(nil))
	return nil
}
