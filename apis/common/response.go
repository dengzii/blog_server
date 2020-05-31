package common

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type BaseController struct {
}

type Json struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func ErrorResponse(status int, msg string, data interface{}) *Json {

	return &Json{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
}

func ErrorEmptyResponse(msg string) *Json {

	return &Json{
		Status: 1000,
		Msg:    msg,
		Data:   nil,
	}
}

func SuccessResponse(data interface{}) (result *Json) {

	result = &Json{
		Status: 200,
		Msg:    "success",
		Data:   data,
	}
	if data == nil {
		result.Msg = "no data."
	}
	return
}

func ReturnServerError(ctx context.Context) {
	msg := ctx.Values().GetString("error")
	result := "500 Internal Server Error"
	if msg == "" {
		result = "Unexpected Internal Server Error"
	}
	_, _ = ctx.Writef(result, msg)
}

func ReturnNotFound(ctx context.Context) {
	ctx.StatusCode(iris.StatusNotFound)
	_, _ = ctx.Writef("Path Not Found", ctx.Path())
}
