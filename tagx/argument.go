package tagx

import "strings"

// DefaultArgHeadSep 默认的tag嵌套参数头分隔符
var DefaultArgHeadSep = ":"

// DefaultArgExtSep 默认的tag嵌套参数体分割符
var DefaultArgExtSep = "|"

// ParseArgs 解析tag嵌套参数的方法
func ParseArgs(str string) (string, []string) {
	parts := strings.SplitN(str, DefaultArgHeadSep, 2)
	if len(parts) == 1 {
		return str, nil
	}
	return parts[0], strings.Split(parts[1], DefaultArgExtSep)
}
