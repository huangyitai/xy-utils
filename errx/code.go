package errx

// Code ...
type Code int

// Error ...
func (c Code) Error(msg string, detail string) *Error {
	return &Error{
		Code:   c,
		Msg:    msg,
		Detail: detail,
		Data:   nil,
	}
}

// ErrorWithData ...
func (c Code) ErrorWithData(msg string, detail string, data interface{}) *Error {
	return &Error{
		Code:   c,
		Msg:    msg,
		Detail: detail,
		Data:   data,
	}
}

// SimpleError ...
func (c Code) SimpleError(msg string) *Error {
	return c.Error(msg, "")
}

// SimpleErrorWithData ...
func (c Code) SimpleErrorWithData(msg string, data interface{}) *Error {
	return c.ErrorWithData(msg, "", data)
}

// ErrorWrapper ...
type ErrorWrapper func(error) *Error

// WrapError ...
func (c Code) WrapError(msg string) ErrorWrapper {
	return func(err error) *Error {
		if err == nil {
			return c.Error(msg, "")
		}

		//如果需要包装的错误已经是errx.Error则不再重复包装
		if e, ok := err.(*Error); ok {
			return e
		}

		return &Error{
			Err:    err,
			Code:   c,
			Msg:    msg,
			Detail: err.Error(),
			Data:   nil,
		}
	}
}

// WrapErrorWithData ...
func (c Code) WrapErrorWithData(msg string, data interface{}) ErrorWrapper {
	return func(err error) *Error {
		if err == nil {
			return c.ErrorWithData(msg, "", data)
		}

		//如果需要包装的错误已经是errx.Error则不再重复包装
		if e, ok := err.(*Error); ok {
			return e
		}

		return &Error{
			Err:    err,
			Code:   c,
			Msg:    msg,
			Detail: err.Error(),
			Data:   data,
		}
	}
}
