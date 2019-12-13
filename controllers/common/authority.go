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

	ctx.Next()
}
