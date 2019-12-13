package user

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"server/controllers"
	"server/models"
)

type loginJson struct {
	Name    string `json:"name"`
	Passwd  string `json:"passwd"`
	Captcha string `json:"captcha"`
}

func LoginController(ctx context.Context) (err error) {

	var requestUser loginJson

	err = ctx.ReadJSON(&requestUser)
	var user = models.GetUser(requestUser.Name, requestUser.Passwd)

	response := controllers.CommonJson(400, "wrong password or username", nil)
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

	var user models.UserJson
	err = ctx.ReadJSON(&user)
	if err != nil {
		return err
	}
	_, err = ctx.JSON(models.AddUser(user.Name, user.Passwd, user.Email))
	return err
}
