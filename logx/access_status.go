package logx

// AccessStatus 请求概述状态，主要区分成功、失败、出错、panic，默认状态为未知
type AccessStatus int

// AccessStatusServerUnknown ...
const (
	AccessStatusServerUnknown AccessStatus = iota
	AccessStatusServerSuccess
	AccessStatusServerFail
	AccessStatusServerError
	AccessStatusServerPanic

	AccessStatusClientUnknown
	AccessStatusClientSuccess
	AccessStatusClientFail
	AccessStatusClientError
	AccessStatusClientPanic
)

// DefaultAccessStatusIcon 默认请求状态图标（业务可以自定义）
var DefaultAccessStatusIcon = map[AccessStatus]string{
	AccessStatusServerUnknown: "⬜",          //白
	AccessStatusServerSuccess: "\U0001F7E9", //绿
	AccessStatusServerFail:    "\U0001F7E8", //黄
	AccessStatusServerError:   "\U0001F7E7", //橙
	AccessStatusServerPanic:   "\U0001F7E5", //红

	AccessStatusClientUnknown: "⚪",          //白
	AccessStatusClientSuccess: "\U0001F7E2", //绿
	AccessStatusClientFail:    "\U0001F7E1", //黄
	AccessStatusClientError:   "\U0001F7E0", //橙
	AccessStatusClientPanic:   "🔴",          //红
}

// GetAccessStatusIcon 根据请求状态获取请求状态图标的方法（业务可以自定义）
var GetAccessStatusIcon = func(status AccessStatus) string {
	return DefaultAccessStatusIcon[status]
}
