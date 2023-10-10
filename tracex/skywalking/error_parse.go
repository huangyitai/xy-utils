package skywalking

import (
	"github.com/SkyAPM/go2sky"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

// RetJSONPath ...
const (
	RetJSONPath    = "ret"
	MsgJSONPath    = "msg"
	DetailJSONPath = "detail"
)

// ParseError 解析响应中包含的错误信息
func ParseError(span go2sky.Span, rsp interface{}) {
	var ret int
	var msg string
	var detail string

	switch rsp.(type) {
	case nil:
		return
	case []byte:
		ret, msg, detail = parseJSONBytes(rsp.([]byte))
	default:
		ret, msg, detail = parseInterface(rsp)
	}

	// 增加状态码上报逻辑，用于 智研APM 上报状态码
	span.Tag(go2sky.TagStatusCode, cast.ToString(ret))
	if isErrorResponse(ret, msg, detail) {
		span.Error(time.Now(),
			RetJSONPath, strconv.Itoa(ret),
			MsgJSONPath, msg,
			DetailJSONPath, detail)
	}
}

func parseJSONBytes(bs []byte) (int, string, string) {
	ret := gjson.GetBytes(bs, RetJSONPath).Int()
	msg := gjson.GetBytes(bs, MsgJSONPath).String()
	detail := gjson.GetBytes(bs, DetailJSONPath).String()
	return int(ret), msg, detail
}

func parseInterface(itf interface{}) (int, string, string) {
	var ret int
	var msg string
	var detail string

	//尝试解析ret
	ra, ok := itf.(retAware)
	if ok {
		ret = int(ra.GetRet())
	}

	//尝试解析msg
	ma, ok := itf.(msgAware)
	if ok {
		msg = ma.GetMsg()
	}

	//尝试解析detail
	da, ok := itf.(detailAware)
	if ok {
		detail = da.GetDetail()
	}

	return ret, msg, detail
}

// retAware 包含ret返回值
type retAware interface {
	GetRet() int32
}

// msgAware 包含msg返回值
type msgAware interface {
	GetMsg() string
}

// detailAware 包含detail返回值
type detailAware interface {
	GetDetail() string
}

func isErrorResponse(ret int, msg, detail string) bool {
	return ret < 0
}
