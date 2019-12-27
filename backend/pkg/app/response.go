package app

import (
	"fmt"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/e"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GenResponse(code int, data interface{}, err error) *Response {
	var msg = ""
	if err != nil {
		msg = fmt.Sprintf("%v: %v", e.GetMsg(code), err)
	} else {
		msg = e.GetMsg(code)
	}
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
