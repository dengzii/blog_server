package main

import (
	"github.com/kataras/iris"
	"server/bootstrap"
	"server/db"
	"server/models"
	"server/routers"
)

var app *bootstrap.Bootstrapper

func main() {

	db.Init()
	models.Init()
	app = bootstrap.New("dengzi's blog", "dengzi", true)

	config := iris.YAML("./conf/iris.yml")
	app.Configure(iris.WithConfiguration(config))
	app.SetupViews("./views")
	app.Bootstrap()

	routers.Setup(app)
	app.Listen(":80")
}
