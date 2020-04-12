package common

import (
	context2 "github.com/kataras/iris/context"
)

var authorityUrl map[string]string

func init() {

}

func AuthorityController(ctx context2.Context) {

	//path := ctx.Path()
	//authorityUrl[path]
	ctx.Header("Vary", "Access-Control-Request-Method")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Request-Headers", "Accept,content-type,X-Requested-With,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,token")
	ctx.Header("Access-Control-Request-Method", "*")
	ctx.Next()
}
