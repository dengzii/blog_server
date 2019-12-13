package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"log"
	"server/bootstrap"
	"server/controllers"
	"server/controllers/article"
	"server/controllers/common"
	"server/controllers/user"
)

func Setup(app *bootstrap.Bootstrapper) {

	app.Use(common.AuthorityController)
	app.Handle("GET", "/", controllers.HomeController)

	app.Handle("GET", "/friends", catchErrorRouter(controllers.GetFriendsController))
	app.Handle("POST", "/friends", catchErrorRouter(controllers.AddFriendsController))
	app.Handle("POST", "/user/login", catchErrorRouter(user.LoginController))
	app.Handle("POST", "/user/register", catchErrorRouter(user.Register))
	app.Handle("POST", "/article/{id}", catchErrorRouter(article.GetArticle))

	app.Handle("GET", "/user/login", catchErrorView("login.html", "", nil))

	app.WildcardSubdomain(subdomainRouter)

	errorRouter(app)
	staticRouter(app)
}

func catchErrorView(view string, dataKey string, dataValue interface{}) func(context.Context) {

	return func(ctx context.Context) {
		if dataKey != "" && dataValue != nil {
			ctx.ViewData(dataKey, dataValue)
		}
		err := ctx.View(view)
		if err != nil {
			_, _ = ctx.WriteString("error," + err.Error())
		}
	}
}

func catchErrorRouter(router func(context.Context) error) func(context.Context) {
	return func(ctx context.Context) {
		err := router(ctx)
		if err != nil {
			_, _ = ctx.WriteString("error," + err.Error())
		}
	}
}

func subdomainRouter(ctx context.Context) {
	ctx.Application().Logger().Info("=> " + ctx.Subdomain())
}

func errorRouter(app *bootstrap.Bootstrapper) {
	app.OnErrorCode(iris.StatusInternalServerError, common.ServerErrorController)
	app.OnErrorCode(iris.StatusNotFound, common.NotFoundController)

}

func staticRouter(app *bootstrap.Bootstrapper) {
	app.Get("/static/{f:string}", func(context context.Context) {
		err := context.ServeFile("./static/a.png", false)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func init() {

}
