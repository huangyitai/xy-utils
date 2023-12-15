package logx

import (
	"context"
	"fmt"

	"github.com/huangyitai/xy-utils/errx"
	"github.com/rs/zerolog"
	"github.com/spf13/cast"
)

// ErrorHandler ...
type ErrorHandler func(ctx context.Context, err error) (*ErrorInfo, bool)

// ErrorHandlerList ...
var ErrorHandlerList = []ErrorHandler{GetErrorInfoFromErrx}

// GetErrorInfoFromErrx ...
func GetErrorInfoFromErrx(ctx context.Context, err error) (*ErrorInfo, bool) {
	if e, ok := err.(*errx.Error); ok {
		return &ErrorInfo{
			Code:          cast.ToString(e.GetCode()),
			Msg:           e.Msg,
			Type:          "errx",
			IsCtxCanceled: ctx.Err() == context.Canceled,
			IsCtxTimeout:  ctx.Err() == context.DeadlineExceeded,
			IsInternal:    e.GetCode() < 0,
		}, true
	}
	return nil, false
}

// GetErrorInfo ...
func GetErrorInfo(ctx context.Context, err error) *ErrorInfo {
	if err == nil {
		return &ErrorInfo{
			Code:          "0",
			Msg:           "success",
			Type:          "success",
			IsCtxCanceled: ctx.Err() == context.Canceled,
			IsCtxTimeout:  ctx.Err() == context.DeadlineExceeded,
			IsInternal:    false,
		}
	}
	for _, handler := range ErrorHandlerList {
		if info, ok := handler(ctx, err); ok {
			return info
		}
	}
	return &ErrorInfo{
		Code:          "unknown",
		Msg:           "",
		Type:          fmt.Sprintf("%T", err),
		IsCtxCanceled: ctx.Err() == context.Canceled,
		IsCtxTimeout:  ctx.Err() == context.DeadlineExceeded,
		IsInternal:    true,
	}
}

// ErrorInfo ...
type ErrorInfo struct {
	Code          string
	Msg           string
	Type          string
	IsCtxTimeout  bool
	IsCtxCanceled bool
	IsInternal    bool
}

// GetInfo 生成错误信息
func (i *ErrorInfo) GetInfo() string {
	return i.Type + ":" + i.Code
}

// MarshalZerologObject ...
func (i *ErrorInfo) MarshalZerologObject(e *zerolog.Event) {
	e.Str("sErrCode", i.Code).Str("sErrMsg", i.Msg).Str("sErrType", i.Type).
		Str("sErrInfo", fmt.Sprintf("%s:%s", i.Type, i.Code)).
		Bool("bCtxTimeout", i.IsCtxTimeout).Bool("bCtxCanceled", i.IsCtxCanceled)
}
