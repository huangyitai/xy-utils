package logx

import (
	"fmt"

	"github.com/huangyitai/xy-utils/xxx"
)

type funcStr struct {
	f func() string
}

// String 调用函数生成字符串
func (j *funcStr) String() string {
	if j == nil {
		return ""
	}
	return xxx.TruncateToMultiLine(j.f(), MaxLenPayloadPrefix, MaxLenPayloadSuffix)
}

// FuncStr 创建一个返回字符串的函数构成的Stringer
func FuncStr(f func() string) fmt.Stringer {
	return &funcStr{f: f}
}
