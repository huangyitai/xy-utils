package errx

import "context"

// LogHandler 错误的日志处理函数
var LogHandler func(ctx context.Context, i interface{}, e *Error)

// Log 快速日志
func (e *Error) Log() *Error {
	return e.logCtxWith(context.Background(), nil)
}

// LogWith 包含自定义信息的日志
func (e *Error) LogWith(i interface{}) *Error {
	return e.logCtxWith(context.Background(), i)
}

// LogCtx 包含上下文信息的快速日志
func (e *Error) LogCtx(ctx context.Context) *Error {
	return e.logCtxWith(ctx, nil)
}

// LogCtxWith 包含上下文信息和自定义信息的日志
func (e *Error) LogCtxWith(ctx context.Context, i interface{}) *Error {
	return e.logCtxWith(ctx, i)
}

func (e *Error) logCtxWith(ctx context.Context, i interface{}) *Error {
	if LogHandler != nil {
		LogHandler(ctx, i, e)
	}
	return e
}
