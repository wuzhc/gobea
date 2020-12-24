package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ErrCode = 0
)

type SysError struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (err *SysError) Error() string {
	return err.Msg
}

func (err *SysError) MarshalJSON() ([]byte, error) {
	return json.Marshal(SysError{
		Msg:  err.Msg,
		Code: err.Code,
		Data: err.Data,
	})
}

func (err *SysError) Headers() http.Header {
	var headers = make(http.Header)
	headers.Add("Content-Type", "application/json")
	return headers
}

func (err *SysError) StatusCode() int {
	return http.StatusOK
}

func NewErr(msg string, code int, data interface{}) *SysError {
	return &SysError{
		Msg:  msg,
		Code: code,
		Data: data,
	}
}

func NewErrMsg(msg string) *SysError {
	return &SysError{
		Msg:  msg,
		Code: ErrCode,
		Data: nil,
	}
}

func NewParamErr(param string) *SysError {
	return &SysError{
		Msg:  fmt.Sprintf("%s参数无效", param),
		Code: ErrCode,
		Data: nil,
	}
}

func NewNotFoundErr() *SysError {
	return &SysError{
		Msg:  "数据不存在或已删除",
		Code: ErrCode,
		Data: nil,
	}
}

func NewNotAccessErr() *SysError {
	return &SysError{
		Msg:  "无权访问",
		Code: ErrCode,
		Data: nil,
	}
}
