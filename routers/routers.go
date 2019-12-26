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

	mainRouter := app.Party("/")

	userRouter := mainRouter.Party("/{username:alphabetical}")

	userRouter.Handle("GET", "/", catchErrorRouter(controllers.HomeController))

	userArticle := userRouter.Party("/articles")
	userArticle.Handle("GET", "/", func(i context.Context) {
		_, _ = i.WriteString("get article from user: " + i.Params().GetString("username"))
	})

	userArticle.Handle("GET", "/", func(i context.Context) {
		_, _ = i.WriteString("get article from user: " + i.Params().GetString("username"))
	})
	userArticle.Handle("GET", "/{id:uint}", func(i context.Context) {
		_, _ = i.WriteString("get article id " + i.Params().GetString("id") + ", from user " + i.Params().GetString("username"))
	})

	app.Handle("GET", "/home", catchErrorRouter(controllers.HomeController))

	app.Handle("GET", "/friends", catchErrorRouter(controllers.GetFriendsController))
	app.Handle("PUT", "/friends", catchErrorRouter(controllers.AddFriendsController))

	app.Handle("GET", "/article/{id:uint}", catchErrorRouter(article.GetArticle))
	app.Handle("PUT", "/article", catchErrorRouter(article.AddArticle))
	app.Handle("GET", "/article", catchErrorRouter(article.GetArticleLatest))

	app.Handle("PUT", "/tag", catchErrorRouter(article.AddTagController))
	app.Handle("GET", "/tag", catchErrorRouter(article.GetTagsController))
	app.Handle("PUT", "/category", catchErrorRouter(controllers.AddCategoryController))

	app.Handle("GET", "/category", catchErrorRouter(controllers.GetCategoriesController))

	app.Handle("GET", "/user/login", catchErrorView("login.html", "", nil))

	usersRouters := app.Party("/user")
	usersRouters.Put("/", catchErrorRouter(user.Register))
	usersRouters.Get("/", catchErrorRouter(user.LoginController))

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
