package errx

import (
	"errors"
	"fmt"
)

// Error ...
type Error struct {
	Err    error
	Code   Code
	Msg    string
	Detail string
	Data   interface{}
}

// Error ...
func (e *Error) Error() string {
	return fmt.Sprintf("error code:%d, msg:%s, detail:%s, data:%+v", e.Code, e.Msg, e.Detail, e.Data)
}

// Is 实现此方法以支持errors.Is判断
func (e *Error) Is(err error) bool {
	if err == nil || e == nil {
		return e == nil && err == nil
	}
	switch t := err.(type) {
	case *Error:
		return e.Code == t.Code
	default:
		return e.Err != nil && errors.Is(e.Err, t)
	}
}

// WithData ...
func (e Error) WithData(data interface{}) *Error {
	e.Data = data
	return &e
}

// WithMsg ...
func (e Error) WithMsg(msg string) *Error {
	e.Msg = msg
	return &e
}

// WithDetail ...
func (e Error) WithDetail(detail string) *Error {
	e.Detail = detail
	return &e
}

// FormatMsg ...
func (e Error) FormatMsg(a ...interface{}) *Error {
	e.Msg = fmt.Sprintf(e.Msg, a...)
	return &e
}

// FormatDetail ...
func (e Error) FormatDetail(a ...interface{}) *Error {
	e.Detail = fmt.Sprintf(e.Detail, a...)
	return &e
}

// Wrap ...
func (e Error) Wrap(err error) *Error {
	if err == nil {
		return &e
	}

	if te, ok := err.(*Error); ok {
		return te
	}

	e.Detail = err.Error()
	e.Err = err
	return &e
}

// Cause ...
func (e Error) Cause(err error) *Error {
	if err == nil {
		return &e
	}

	e.Detail = err.Error()
	e.Err = err
	return &e
}

// GetCode 获取错误code
func (e *Error) GetCode() int {
	if e == nil {
		return 0
	}
	return int(e.Code)
}
