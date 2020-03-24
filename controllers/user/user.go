package user

import (
	"github.com/dengzii/blog_server/controllers"
	user2 "github.com/dengzii/blog_server/models/user"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type loginJson struct {
	Name    string `json:"name"`
	Passwd  string `json:"passwd"`
	Captcha string `json:"captcha"`
}

func LoginController(ctx context.Context) (err error) {

	var requestUser loginJson

	err = ctx.ReadJSON(&requestUser)
	var user = user2.GetUser(requestUser.Name, requestUser.Passwd)

	response := controllers.ErrorResponse(400, "wrong password or username", nil)
	if user.Name != "" {
		response.Status = 200
		response.Msg = "success"
		response.Data = user
	}
	ctx.StatusCode(iris.StatusOK)
	_, err = ctx.JSON(response)
	return err
}

func Register(ctx context.Context) (err error) {

	var user user2.UserJson
	err = ctx.ReadJSON(&user)
	if err != nil {
		return err
	}
	_, err = ctx.JSON(user2.AddUser(user.Name, user.Passwd, user.Email))
	return err
}
