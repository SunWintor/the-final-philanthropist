package ecode

import (
	"errors"
)

type Ecode struct {
	Code int64
	Msg  string
	error
}

func newErr(code int64, msg, err string) *Ecode {
	return &Ecode{Code: code, Msg: msg, error: errors.New(err)}
}

var (
	UnknownError        = newErr(-1, "未知错误", "unkonwn")
	NotLoginErr         = newErr(-101, "未登录", "not login")
	PermissionDeniedErr = newErr(-401, "无权访问", "permission denied")
)

func EqualError(code *Ecode, err error) bool {
	if e, ok := err.(Ecode); ok {
		return e.Code == code.Code
	}
	return false
}

func ToEcode(err error) (e *Ecode, ok bool) {
	e, ok = err.(*Ecode)
	return
}
