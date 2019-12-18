package controllers

import "github.com/kataras/iris/context"

const (
	ERR_CODE_UPDATE_DB  = 1
	ERR_CODE_LOSE_PARAM = 2
	ERR_CODE_UNEXPECTED = 3
	ERR_CODE_PARSE      = 4
)

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
func NewControllerError(msg string, code int) error {
	return ControllerErrorIm{
		ErrMsg:    msg,
		ErrCode:   code,
		ErrOrigin: nil,
	}
}
