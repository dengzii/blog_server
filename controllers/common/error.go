package common

import "github.com/kataras/iris/context"

func ServerErrorController(ctx context.Context) {
	msg := ctx.Values().GetString("error")
	result := "<h1 style='color:red;'>Internal Server Error : %s</h1>"
	if msg == "" {
		result = "<h1 style='color:red;'>Unexcepted Internal Server Error%s</h1>"
	}
	_, _ = ctx.Writef(result, msg)
}

func NotFoundController(ctx context.Context) {
	_, _ = ctx.Writef("<h1 style='color:red;'>Not Found Path  %s</h1>", ctx.Path())
}
