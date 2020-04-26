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
	ctx.Header("Allow", "POST, GET, PUT, DELETE, OPTIONS")
	ctx.Header("Vary", "Access-Control-Request-Method")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Request-Headers", "Accept,X-Requested-With,Content-Length, Accept-Encoding,X-CSRF-Token,Authorization,token, Content-Type")
	ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type")
	ctx.Header("Access-Control-Request-Method", "*")
	ctx.Next()
}
