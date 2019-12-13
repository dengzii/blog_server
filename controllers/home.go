package controllers

import "github.com/kataras/iris/context"

func HomeController(ctx context.Context) {
	ctx.ViewData("username", ctx.Subdomain())
	_ = ctx.View("index.html")
}
