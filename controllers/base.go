package controllers

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

func SuccessResponse(data interface{}) *Json {

	return &Json{
		Status: 200,
		Msg:    "success",
		Data:   data,
	}
}
