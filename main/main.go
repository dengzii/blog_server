package main

import (
	"fmt"
	"github.com/dengzii/blog_server/bootstrap"
	"github.com/dengzii/blog_server/conf"
	"github.com/dengzii/blog_server/db"
	"github.com/dengzii/blog_server/models"
	"github.com/dengzii/blog_server/routers"
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
