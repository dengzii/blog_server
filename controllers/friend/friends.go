package friend

import (
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/friend"
	"github.com/kataras/iris/context"
)

func AddFriendsController(ctx context.Context) (err error) {

	var f *friend.Friend
	err = ctx.ReadJSON(&f)
	result := friend.AddFriend(f)
	if result != nil {
		_, _ = ctx.JSON(controllers.SuccessResponse(nil))
	} else {
		_, _ = ctx.JSON(controllers.ErrorResponse(500, "Something went wrong.", nil))
	}
	return err
}

func GetFriendsController(ctx context.Context) (err error) {

	friends := friend.GetFriend()
	_, err = ctx.JSON(controllers.SuccessResponse(friends))
	return err
}
