package friend

import (
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/friend"
	"github.com/kataras/iris/context"
)

func AddFriendsController(ctx context.Context) (err error) {

	var f *friend.Friend
	err = ctx.ReadJSON(&f)
	_ = friend.AddFriend(f)
	_, _ = ctx.JSON(controllers.SuccessResponse(nil))
	return err
}

func GetFriendsController(ctx context.Context) (err error) {

	friends := friend.GetFriend()
	_, err = ctx.JSON(controllers.SuccessResponse(friends))
	return err
}
