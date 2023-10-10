package xxx

import (
	"strconv"
	"strings"
)

// UnquoteUnicode 去除字符串中utf8编码的unicode转义 \uXXXX
func UnquoteUnicode(raw string) (string, error) {
	str, err := strconv.Unquote(
		strings.Replace(
			strconv.Quote(raw), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}
