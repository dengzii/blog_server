package main

import (
	"server/bootstrap"
	"server/db"
	"server/routers"
)

var app *bootstrap.Bootstrapper

func main() {

	db.Init()
	app = bootstrap.New("dengzi's blog", "dengzi", true)

	app.SetupViews("./views")
	app.Bootstrap()

	routers.Setup(app)
	app.Listen(":80")
}
