package errx

import "github.com/pkg/errors"

// NewError ...
func NewError(code int, msg string, detail string) *Error {
	return Code(code).Error(msg, detail)
}

// NewErrorWithData ...
func NewErrorWithData(code int, msg string, detail string, data interface{}) *Error {
	return Code(code).ErrorWithData(msg, detail, data)
}

// NewSimpleError ...
func NewSimpleError(code int, msg string) *Error {
	return Code(code).SimpleError(msg)
}

// NewSimpleErrorWithData ...
func NewSimpleErrorWithData(code int, msg string, data interface{}) *Error {
	return Code(code).SimpleErrorWithData(msg, data)
}

// NewWrapError ...
func NewWrapError(code int, msg string) ErrorWrapper {
	return Code(code).WrapError(msg)
}

// NewWrapErrorWithData ...
func NewWrapErrorWithData(code int, msg string, data interface{}) ErrorWrapper {
	return Code(code).WrapErrorWithData(msg, data)
}

// UnknownCode ...
const UnknownCode Code = -1_234_567

// GetCode ...
func GetCode(err error) Code {
	err = errors.Cause(err)
	if err == nil {
		return 0
	}

	if e, ok := err.(*Error); ok {
		return e.Code
	} else {
		return UnknownCode
	}
}

// GetMsg ...
func GetMsg(err error) string {
	err = errors.Cause(err)
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok {
		return e.Msg
	} else {
		return err.Error()
	}
}

// GetDetail ...
func GetDetail(err error) string {
	err = errors.Cause(err)
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok {
		return e.Detail
	} else {
		return err.Error()
	}
}

// GetData ...
func GetData(err error) interface{} {
	err = errors.Cause(err)
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		return e.Data
	} else {
		return nil
	}
}

// GetErr ...
func GetErr(err error) error {
	err = errors.Cause(err)
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		return e.Err
	} else {
		return err
	}
}

// GetError ...
func GetError(err error) *Error {
	err = errors.Cause(err)
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		return e
	} else {
		return &Error{
			Err:    err,
			Code:   UnknownCode,
			Msg:    err.Error(),
			Detail: err.Error(),
			Data:   nil,
		}
	}
}
