package bootstrap

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"time"
)

type Configuretor func(*Bootstrapper)

var debugMode bool

const (
	StaticAssets = "./statics/"
	Favicon      = "dengzi/favicon.png"
	Views        = "./views"
	ViewEx       = ".html"
)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnData time.Time

	Sessions *sessions.Sessions
}

func New(appName, appOwner string, debug bool, cfgs ...Configuretor) *Bootstrapper {

	debugMode = debug
	b := &Bootstrapper{
		Application:  iris.New(),
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnData: time.Now(),
	}
	b.Application.Configure(iris.WithCharset("UTF-8"))
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

//func (b *Bootstrapper) RegisterMiddleware(middleware middleware.Middleware) {
//	middleware.Attach(b)
//}

func (b *Bootstrapper) SetupViews(viewDir string) {

	templateEngine := iris.HTML(viewDir, ViewEx).Reload(debugMode)
	b.RegisterView(templateEngine)
}

func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:  "SECRET_SESS_COOKIE_" + b.AppName,
		Expires: expires,
	})
}

func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(context context.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  context.GetStatusCode(),
			"message": context.Values().GetString("message"),
		}
		if json := context.URLParamExists("json"); json {
			context.JSON(err)
			return
		}
		context.ViewData("error", err)
		context.View("error.html")
	})
}

func (b *Bootstrapper) Confiture(cs ...Configuretor) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews(Views)
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()

	// static files
	//b.StaticHandler(StaticAssets, false, false)
	b.Favicon(StaticAssets + Favicon)
	//b.HandleDir(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	b.Use(recover2.New())
	b.Use(logger.New())

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
