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
	app = bootstrap.New("dengzi's blog", "dengzi", true)

	fmt.Println(conf.Get())
	app.SetupViews("./views")
	app.Bootstrap()

	routers.Setup(app)
	app.Listen(":80")
}
