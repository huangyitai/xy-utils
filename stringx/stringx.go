package stringx

import "strings"

// SplitAndTrimSpace ...
func SplitAndTrimSpace(str string, sep string) []string {
	res := strings.Split(str, sep)
	for i, s := range res {
		res[i] = strings.TrimSpace(s)
	}
	return res
}
