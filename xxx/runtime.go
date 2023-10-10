package xxx

import (
	"fmt"
	"runtime"
)

// FuncCaller 返回当前函数（调用FuncCaller的函数）的上层调用函数的代码行地址
func FuncCaller() string {
	return Caller(2)
}

// Caller 返回当前函数（调用FuncCaller的函数为skip=0）被调的上层代码行地址
func Caller(skip int) string {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}
