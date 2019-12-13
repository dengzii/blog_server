package controllers

type BaseController struct {
}

type Json struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func CommonJson(status int, msg string, data interface{}) *Json {

	return &Json{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
}
