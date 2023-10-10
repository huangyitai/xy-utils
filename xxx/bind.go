package xxx

import (
	"github.com/huangyitai/xy-utils/bindx"
	"github.com/rs/zerolog/log"
)

// StartBind 开始绑定前初始化日志签名，打印trace日志的帮助函数
func StartBind(slot bindx.Slot, binderDirs ...string) *SignStr {
	sign := NewSignStr().WithPath(binderDirs...).WithProp(bindx.TagKey, slot.Name())
	log.Trace().Str("sName", slot.Name()).Str("sTag", string(slot.Tag())).Msgf("%s bind start", sign)
	return sign
}
