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

	app.Handle("GET", "/home", catchErrorRouter(controllers.HomeController))

	mainRouter := app.Party("/")

	friendRouter := mainRouter.Party("/friend")
	friendRouter.Get("/", catchErrorRouter(controllers.GetFriendsController))
	friendRouter.Put("/", catchErrorRouter(controllers.AddFriendsController))
	friendRouter.Patch("/", catchErrorRouter(controllers.AddFriendsController))

	articleRouter := mainRouter.Party("/article")
	articleRouter.Get("/{id:uint}", catchErrorRouter(article.GetArticle))
	articleRouter.Put("/", catchErrorRouter(article.AddArticle))
	articleRouter.Get("/latest", catchErrorRouter(article.GetArticleLatest))

	tagRouter := mainRouter.Party("tag")
	tagRouter.Put("/", catchErrorRouter(article.AddTagController))
	tagRouter.Get("/", catchErrorRouter(article.GetTagsController))
	//tagRouter.Patch("/", catchErrorRouter(article.UpdateTagsController))

	categoryRouter := mainRouter.Party("category")
	categoryRouter.Get("/", catchErrorRouter(controllers.AddCategoryController))
	categoryRouter.Put("/", catchErrorRouter(controllers.GetCategoriesController))
	//categoryRouter.Patch("/", catchErrorRouter(controllers.PatchCategoriesController))

	userRouter := app.Party("/user")
	userRouter.Put("/", catchErrorRouter(user.Register))
	userRouter.Post("/", catchErrorRouter(user.LoginController))
	//userRouter.Get("", catchErrorRouter(user.ProfileController))
	//userRouter.Patch("/", catchErrorRouter(user.UpdateController))

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
