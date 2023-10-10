package poolx

import (
	"bytes"
	"sync"
)

var bytesBufferPool = &sync.Pool{New: func() interface{} { return &bytes.Buffer{} }}

// GetBytesBuffer 从池中获取buffer对象
func GetBytesBuffer() *bytes.Buffer {
	return bytesBufferPool.Get().(*bytes.Buffer)
}

// PutBytesBuffer 向池中放回buffer对象
func PutBytesBuffer(buf *bytes.Buffer) {
	if buf != nil {
		buf.Reset()
		bytesBufferPool.Put(buf)
	}
}

// UseBytesBuffer 使用buffer对象执行逻辑，结束后buffer将会被自动放回
// 请勿在函数 f 外使用或持有buffer对象，否则将会出现数据竞争等异常
func UseBytesBuffer(f func(buf *bytes.Buffer)) {
	buf := GetBytesBuffer()
	defer PutBytesBuffer(buf)
	f(buf)
}
