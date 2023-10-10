package xxx

import (
	"fmt"
	"strings"
)

// SignStr 签名字符串
type SignStr struct {
	strings.Builder
}

// NewSignStr ...
func NewSignStr() *SignStr {
	return &SignStr{}
}

// WithPath 添加一个path格式的签名部分，如[a/b/c/d]
func (s *SignStr) WithPath(dirs ...string) *SignStr {
	if len(dirs) == 0 {
		_, _ = s.WriteString("[]")
		return s
	}

	_, _ = s.WriteString("[" + dirs[0])

	for _, dir := range dirs[1:] {
		_, _ = s.WriteString("/" + dir)
	}
	_, _ = s.WriteString("]")
	return s
}

// WithProp 添加一个prop格式的签名部分，如[key:value]
func (s *SignStr) WithProp(key, value string) *SignStr {
	_, _ = s.WriteString("[" + key)
	_, _ = s.WriteString(":" + value + "]")
	return s
}

// Sign 为一条消息签名
func (s *SignStr) Sign(msg string) string {
	if s.Len() == 0 {
		return msg
	}
	return s.String() + " " + msg
}

// Signf 为一条格式消息签名
func (s *SignStr) Signf(format string, args ...interface{}) string {
	if s.Len() == 0 {
		return fmt.Sprintf(format, args...)
	}
	return s.String() + " " + fmt.Sprintf(format, args...)
}
