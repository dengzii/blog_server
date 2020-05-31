package routers

import (
	"fmt"
	"github.com/dengzii/blog_server/apis/article"
	"github.com/dengzii/blog_server/apis/common"
	"github.com/dengzii/blog_server/apis/friend"
	"github.com/dengzii/blog_server/apis/user"
	"github.com/dengzii/blog_server/bootstrap"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
)

func Setup(app *bootstrap.Bootstrapper) {

	app.Use(common.AuthorityController)

	mainRouter := app.Party("/{username:string  regexp(^[A-Za-z0-9]{4,12})}")

	mainRouter.Get("/", catchErrorRouter(article.GetArticlesApi))
	mainRouter.Get("/articles", catchErrorRouter(article.GetArticlesApi))
	mainRouter.Get("/archive", catchErrorRouter(article.GetArchiveApi))
	mainRouter.Get("/profile", catchErrorRouter(user.GetUserInfoApi))

	mainRouter.PartyFunc("/about", aboutRouterFunc)
	mainRouter.PartyFunc("/friend", friendRouterFunc)
	mainRouter.PartyFunc("/article", articleRouterFunc)
	mainRouter.PartyFunc("/category", categoryRouterFunc)
	mainRouter.PartyFunc("/tag", tagRouterFunc)
	mainRouter.PartyFunc("/user", userRouterFunc)

	app.WildcardSubdomain(subdomainRouter)

	errorRouter(app)
	mainRouter.PartyFunc("/static", userStaticRouterFunc)
}

func userRouterFunc(p router.Party) {

	p.Put("/", catchErrorRouter(user.RegisterApi))
	p.Post("/", catchErrorRouter(user.LoginApi))
	//p.Patch("/", catchErrorRouter(user.UpdateController))
}

func aboutRouterFunc(p router.Party) {
	p.Get("/", catchErrorRouter(user.GetAbout))
	p.Put("/", catchErrorRouter(user.AddAbout))
}

func friendRouterFunc(p router.Party) {
	p.Get("/", catchErrorRouter(friend.GetFriendsApi))
	p.Put("/", catchErrorRouter(friend.AddFriendsApi))
	p.Patch("/", catchErrorRouter(friend.AddFriendsApi))
}

func articleRouterFunc(p router.Party) {
	p.Get("/{id:uint}", catchErrorRouter(article.GetArticleApi))
	p.Put("/", catchErrorRouter(article.AddArticleApi))
	p.Get("/", catchErrorRouter(article.GetArticlesApi))
	p.Post("/", catchErrorRouter(article.ViewArticleApi))
}

func tagRouterFunc(p router.Party) {

	p.Put("/", catchErrorRouter(article.AddTagApi))
	p.Get("/", catchErrorRouter(article.GetTagsApi))
	//p.Patch("/", catchErrorRouter(article.UpdateTagsController))
}

func categoryRouterFunc(p router.Party) {

	p.Get("/", catchErrorRouter(article.GetCategoriesApi))
	p.Put("/", catchErrorRouter(article.AddCategoryApi))
	//p.Patch("/", catchErrorRouter(apis.PatchCategoriesController))
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

}

func errorRouter(app *bootstrap.Bootstrapper) {
	app.OnErrorCode(iris.StatusInternalServerError, common.ReturnServerError)
	app.OnErrorCode(iris.StatusNotFound, common.ReturnNotFound)
}

func userStaticRouterFunc(p router.Party) {
	p.Get("/{file:string}", func(context context.Context) {
		f := context.URLParam("file")
		u := context.URLParam("username")
		path := fmt.Sprintf("./static/%s/%s", u, f)
		err := context.ServeFile(path, false)
		if err != nil {
			common.ReturnNotFound(context)
		}
	})
}

func init() {

}
