package user

import (
	"github.com/dengzii/blog_server/apis/common"
	user2 "github.com/dengzii/blog_server/models/user"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type loginJson struct {
	Name    string `json:"name"`
	Passwd  string `json:"passwd"`
	Captcha string `json:"captcha"`
}

type registerJson struct {
	Name    string `json:"name"`
	Passwd  string `json:"passwd"`
	Email   string `json:"email"`
	Captcha string `json:"captcha"`
}

type userIdJson struct {
	Name string `json:"name"`
}

func LoginApi(ctx context.Context) (err error) {

	var requestUser loginJson

	err = ctx.ReadJSON(&requestUser)
	var user = user2.GetUser(requestUser.Name, requestUser.Passwd)

	response := common.ErrorResponse(400, "wrong password or username", nil)
	if user.Name != "" {
		response.Status = 200
		response.Msg = "success"
		response.Data = user
	}
	ctx.StatusCode(iris.StatusOK)
	_, err = ctx.JSON(response)
	return err
}

func RegisterApi(ctx context.Context) (err error) {

	var user registerJson
	err = ctx.ReadJSON(&user)
	if err != nil {
		return err
	}
	success := user2.AddUser(user.Name, user.Passwd, user.Email)
	if success {
		_, err = ctx.JSON(common.SuccessResponse(nil))
	} else {
		_, err = ctx.JSON(common.ErrorEmptyResponse("register failed."))
	}
	return err
}

func GetUserInfoApi(ctx context.Context) (err error) {
	name := ctx.URLParam("user")
	profile := user2.GetUserInfo(name)
	_, err = ctx.JSON(common.SuccessResponse(profile))
	return err
}
