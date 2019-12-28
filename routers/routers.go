package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	"log"
	"server/bootstrap"
	"server/controllers"
	"server/controllers/article"
	"server/controllers/category"
	"server/controllers/common"
	"server/controllers/friend"
	"server/controllers/user"
)

func Setup(app *bootstrap.Bootstrapper) {

	app.Use(common.AuthorityController)

	app.Handle("GET", "/home", catchErrorRouter(controllers.HomeController))

	mainRouter := app.Party("/{username:string  regexp(^[A-Za-z0-9]{4,12})}")

	mainRouter.Get("/", catchErrorRouter(article.GetArticleLatest))

	mainRouter.PartyFunc("/friend", friendRouterFunc)
	mainRouter.PartyFunc("/article", articleRouterFunc)
	mainRouter.PartyFunc("/category", categoryRouterFunc)
	mainRouter.PartyFunc("/tag", tagRouterFunc)
	mainRouter.PartyFunc("/user", userRouterFunc)

	app.WildcardSubdomain(subdomainRouter)

	errorRouter(app)
	staticRouter(app)
}

func userRouterFunc(p router.Party) {

	p.Put("/", catchErrorRouter(user.Register))
	p.Post("/", catchErrorRouter(user.LoginController))
	//p.Get("/", catchErrorRouter(user.ProfileController))
	//p.Patch("/", catchErrorRouter(user.UpdateController))
}

func friendRouterFunc(p router.Party) {
	p.Get("/", catchErrorRouter(friend.GetFriendsController))
	p.Put("/", catchErrorRouter(friend.AddFriendsController))
	p.Patch("/", catchErrorRouter(friend.AddFriendsController))
}

func articleRouterFunc(p router.Party) {
	p.Get("/{id:uint}", catchErrorRouter(article.GetArticle))
	p.Put("/", catchErrorRouter(article.AddArticle))
	p.Get("/latest", catchErrorRouter(article.GetArticleLatest))
}

func tagRouterFunc(p router.Party) {

	p.Put("/", catchErrorRouter(article.AddTagController))
	p.Get("/", catchErrorRouter(article.GetTagsController))
	//p.Patch("/", catchErrorRouter(article.UpdateTagsController))
}

func categoryRouterFunc(p router.Party) {

	p.Get("/", catchErrorRouter(category.AddCategoryController))
	p.Put("/", catchErrorRouter(category.GetCategoriesController))
	//p.Patch("/", catchErrorRouter(controllers.PatchCategoriesController))
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
