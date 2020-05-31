package friend

import (
	"github.com/dengzii/blog_server/apis/common"
	"github.com/dengzii/blog_server/models/friend"
	"github.com/kataras/iris/context"
)

func AddFriendsApi(ctx context.Context) (err error) {

	var f *friend.Friend
	err = ctx.ReadJSON(&f)
	result := friend.AddFriend(f)
	if result != nil {
		_, _ = ctx.JSON(common.SuccessResponse(nil))
	} else {
		_, _ = ctx.JSON(common.ErrorResponse(500, "Something went wrong.", nil))
	}
	return err
}

func GetFriendsApi(ctx context.Context) (err error) {

	friends := friend.GetFriend()
	_, err = ctx.JSON(common.SuccessResponse(friends))
	return err
}
