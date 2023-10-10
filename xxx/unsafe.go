package xxx

import "unsafe"

// UnsafeToString 无拷贝将byte切片转换为string
// 请勿继续修改byte切片内容，以免触发未定义行为
func UnsafeToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// UnsafeToBytes 无拷贝将string转换为byte切片
// 请勿修改返回byte切片内容，以免触发未定义行为
func UnsafeToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
