package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/hero"
	"log"
	"reflect"
	"server/bootstrap"
	"server/controllers"
	"server/controllers/article"
	"server/controllers/common"
	"server/controllers/user"
	"strings"
)

func Setup(app *bootstrap.Bootstrapper) {

	app.Use(common.AuthorityController)

	mainRouter := app.Party("/")

	macro := func(paramValue string) (i interface{}, b bool) {
		return strings.Split(paramValue, "/"), true
	}

	app.Macros().Register("username", "",
		false, true, macro).
		RegisterFunc("contains", func(expectedItem []string) func(paramValue []string) bool {
			return func(paramValue []string) bool {
				return true
			}
		})

	context.ParamResolvers[reflect.TypeOf([]string{})] = func(paramIndex int) interface{} {
		return func(ctx iris.Context) []string {
			return ctx.Params().GetEntryAt(paramIndex).ValueRaw.([]string)
		}
	}

	app.Handle("GET", "/test/{a:username}", hero.Handler(func(param []string) string {
		return "params, " + param[0] + param[1]
	}))

	app.Get("/test/{myparam:username contains([value1,value2])}", func(ctx iris.Context) {
		// When it is not a builtin function available to retrieve your value with the type you want, such as ctx.Params().GetInt
		// then you can use the `GetEntry.ValueRaw` to get the real value, which is set-ed by your macro above.
		myparam := ctx.Params().GetEntry("myparam").ValueRaw.([]string)
		ctx.Writef("myparam's value (a trailing path parameter type) is: %#v\n", myparam)
	})

	userRouter := mainRouter.Party("/{username:alphabetical}")

	userRouter.Handle("GET", "/", catchErrorRouter(controllers.HomeController))

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
