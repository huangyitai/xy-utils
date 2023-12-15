package logx

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Writer 一个zerolog的writer适配器，和zerolog.Logger区别在于，如果未指定logger，每次会采用最新的全局Logger输出日志
type Writer struct {
	l *zerolog.Logger
}

// NewWriter 创建一个新的writer
func NewWriter(l *zerolog.Logger) *Writer {
	return &Writer{l: l}
}

// Write 输出日志方法
func (w *Writer) Write(p []byte) (n int, err error) {
	lg := w.l
	if lg == nil {
		lg = &log.Logger
	}
	return lg.Write(p)
}
