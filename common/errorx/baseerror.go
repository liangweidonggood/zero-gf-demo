package errorx

import (
	"net/http"
)

const defaultCode = 1001

var (
	NoDataErr = &CodeError{Code: 404, Msg: "no data"}
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ErrorHandler(err error) (int, interface{}) {
	return http.StatusConflict, CodeError{
		Code: -1,
		Msg:  err.Error(),
	}
}
func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
