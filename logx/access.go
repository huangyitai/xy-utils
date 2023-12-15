package logx

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/huangyitai/xy-utils/xxx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// 输出payload长度控制：超出前缀+后缀长度的payload，中间部分将会被省略
var (
	// MaxLenPayloadPrefix 输出日志中payload最大前缀长度
	MaxLenPayloadPrefix = 10 * 1024
	// MaxLenPayloadSuffix 输出日志中payload最大后缀长度
	MaxLenPayloadSuffix = 10 * 1024

	// MaxLenNamePrefix Access名称最大前缀长度
	MaxLenNamePrefix = 512
	// MaxLenNameSuffix Access名称最大后缀长度
	MaxLenNameSuffix = 512

	// MaxLenResultPrefix Access结果最大前缀长度
	MaxLenResultPrefix = 16
	// MaxLenResultSuffix Access结果最大后缀长度
	MaxLenResultSuffix = 16
)

// Access 标准Access日志对象
type Access struct {
	ctx      context.Context
	LogEvent *LogEvent
	Sign     *xxx.SignStr
	Status   AccessStatus
	Name     string
	Result   string
	Cost     time.Duration
	Timeout  time.Duration
	Peer     string
	Tags     []string
	Request  fmt.Stringer
	Response fmt.Stringer
	Error    *error
}

// NewAccess 创建新的Access日志对象
func NewAccess(sign string) *Access {
	return &Access{Sign: xxx.NewSignStr().WithPath(sign)}
}

// NewAccessWithEvent 创建新的Access日志对象，包含日志event
func NewAccessWithEvent(sign string, event *zerolog.Event) *Access {
	return &Access{
		Sign:     xxx.NewSignStr().WithPath(sign),
		LogEvent: &LogEvent{event: event},
	}
}

// SetContext 设置Access的ctx信息
func (a *Access) SetContext(ctx context.Context) {
	a.ctx = ctx
}

// Context 返回Access的ctx信息，默认为background
func (a *Access) Context() context.Context {
	if a.ctx == nil {
		return context.Background()
	}
	return a.ctx
}

// LogWithEvent 基于指定event输出access日志，只渲染msg
func (a *Access) LogWithEvent(e *zerolog.Event) {
	if a.Request != nil {
		e.Stringer("sReq", a.Request)
	}
	if a.Response != nil {
		e.Stringer("sRsp", a.Response)
	}
	if a.Error != nil {
		e.Stringer("sErr", JSONStr(*a.Error))
		e.Err(*a.Error)
		e.EmbedObject(GetErrorInfo(a.Context(), *a.Error))
	} else {
		e.EmbedObject(GetErrorInfo(a.Context(), nil))
	}

	tags := ""
	if len(a.Tags) > 0 {
		tags = " | " + strings.Join(a.Tags, " ")
	}
	e.Msgf("%s %s %s | <%s> | %s (%s) | %s%s",
		a.Sign, GetAccessStatusIcon(a.Status), xxx.TruncateToSingleLine(a.Name, MaxLenNamePrefix, MaxLenNameSuffix),
		xxx.TruncateToSingleLine(a.Result, MaxLenResultPrefix, MaxLenResultSuffix), a.Cost, a.Timeout, a.Peer, tags)
}

// FullLogWithEvent 基于指定event输出access日志，包含完整字段信息
func (a *Access) FullLogWithEvent(e *zerolog.Event) {
	e.Stringer("sAccessSign", a.Sign).
		Str("sAccessStatus", GetAccessStatusIcon(a.Status)).
		Str("sAccessName", xxx.TruncateToSingleLine(a.Name, MaxLenNamePrefix, MaxLenNameSuffix)).
		Str("sAccessResult", xxx.TruncateToSingleLine(a.Result, MaxLenResultPrefix, MaxLenResultSuffix)).
		Int64("iAccessCostMicros", a.Cost.Microseconds()).
		Int64("iAccessTimeoutMicros", a.Timeout.Microseconds()).
		Str("sAccessPeer", a.Peer).
		Stringer("sAccessTags", JSONStr(a.Tags))

	a.LogWithEvent(e)
}

// Log 基于Access中包含event输出日志，如果event为空，则创建默认的info日志event
func (a *Access) Log() {
	if a.LogEvent == nil {
		a.LogWithEvent(log.Info())
	}
	a.LogWithEvent(a.LogEvent.event)
}

// FullLog 基于Access中包含event输出日志，如果event为空，则创建默认的info日志event
func (a *Access) FullLog() {
	if a.LogEvent == nil {
		a.FullLogWithEvent(log.Info())
	}
	a.FullLogWithEvent(a.LogEvent.event)
}

// AddTag 向Access添加tag
func (a *Access) AddTag(tags ...string) *Access {
	a.Tags = append(a.Tags, tags...)
	return a
}

// SetReq 向Access添加请求体
func (a *Access) SetReq(req fmt.Stringer) *Access {
	a.Request = req
	return a
}

// SetRsp 向Access添加响应体
func (a *Access) SetRsp(rsp fmt.Stringer) *Access {
	a.Response = rsp
	return a
}

// SetErr 向Access添加错误信息
func (a *Access) SetErr(err error) *Access {
	a.Error = &err
	return a
}
