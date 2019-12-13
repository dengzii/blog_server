package controllers

import "github.com/kataras/iris/context"

type ControllerError interface {
	Error() string
	Code() int
	OriginError() error
	Context() context.Context
}

type ControllerErrorIm struct {
	ErrMsg    string
	ErrCode   int
	ErrOrigin error
}

func (er ControllerErrorIm) Error() string {
	return er.ErrMsg
}

func (er ControllerErrorIm) Code() int {
	return er.ErrCode
}

func (er ControllerErrorIm) OriginError() error {
	return er.ErrOrigin
}
