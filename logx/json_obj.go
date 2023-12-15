package logx

import "fmt"

type jsonStr struct {
	value interface{}
}

// String 将value序列化为json
func (j *jsonStr) String() string {
	if j == nil {
		return ""
	}
	return ToJSONStr(j.value)
}

// JSONStr 生成json字符串序列化对象
func JSONStr(v interface{}) fmt.Stringer {
	return &jsonStr{value: v}
}
