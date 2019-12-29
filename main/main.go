package main

import (
	"fmt"
	"server/bootstrap"
	"server/conf"
	"server/db"
	"server/models"
	"server/routers"
)

var app *bootstrap.Bootstrapper

func main() {

	db.Init()
	models.Init()
	app = bootstrap.New(
		conf.Get().Iris.AppName,
		conf.Get().Iris.Owner,
		true)

	app.SetupViews("./views")
	app.Bootstrap()

	routers.Setup(app)
	app.Listen(fmt.Sprintf(":%d", conf.Get().Iris.Port))
}
