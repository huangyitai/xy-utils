package logx

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"strconv"
	"sync"
	"sync/atomic"
)

var lock sync.Mutex
var gWriters []io.Writer

var offsetFieldName = "@offset"
var offset uint64 = 0

// LevelWriteCloser ...
type LevelWriteCloser interface {
	zerolog.LevelWriter
	io.Closer
}

// LevelWriteCloserWrapper ...
type LevelWriteCloserWrapper struct {
	io.WriteCloser
	Level zerolog.Level
}

// WrapWithLevel ...
func WrapWithLevel(w io.WriteCloser, level zerolog.Level) *LevelWriteCloserWrapper {
	return &LevelWriteCloserWrapper{
		WriteCloser: w,
		Level:       level,
	}
}

// WriteLevel ...
func (l *LevelWriteCloserWrapper) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= l.Level {
		return l.Write(p)
	}
	return len(p), nil
}

func setupFormat(cfg *FormatConfig) {
	fixFormatConfig(cfg)

	zerolog.TimestampFieldName = cfg.TimeKey
	zerolog.TimeFieldFormat = cfg.TimeFormat

	zerolog.LevelFieldName = cfg.LevelKey
	zerolog.MessageFieldName = cfg.MessageKey
	zerolog.ErrorFieldName = cfg.ErrorKey

	zerolog.CallerFieldName = cfg.CallerKey
	zerolog.CallerMarshalFunc = callerMarshal

	zerolog.ErrorStackFieldName = cfg.StacktraceKey
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	offsetFieldName = cfg.OffsetKey
}

func callerMarshal(pc uintptr, file string, line int) string {
	cnt := 0
	pos := len(file) - 1
	for ; pos >= 0; pos-- {
		if file[pos] == '/' {
			cnt++
			if cnt == 2 {
				break
			}
		}
	}
	return file[pos+1:] + ":" + strconv.Itoa(line)
}

func setupWriters(cfg []WriterConfig) ([]io.Writer, error) {
	var writers []io.Writer
	for _, wcfg := range cfg {
		w, err := wcfg.Setup()
		if err != nil {
			return nil, err
		}
		writers = append(writers, w)
	}
	return writers, nil
}

func setupGlobal(cfg *GlobalConfig, w io.Writer) zerolog.Logger {
	logCtx := zerolog.New(w).With()

	if !cfg.DisableCaller {
		logCtx = logCtx.Caller()
	}
	if !cfg.DisableErrorStack {
		logCtx = logCtx.Stack()
	}
	if !cfg.DisableTimestamp {
		logCtx = logCtx.Timestamp()
	}
	if cfg.CustomFields != nil {
		for key, value := range cfg.CustomFields {
			logCtx = logCtx.Str(key, value)
		}
	}

	logger := logCtx.Logger()

	if !cfg.DisableOffset {
		logger = logger.Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
			e.Uint64(offsetFieldName, atomic.AddUint64(&offset, 1))
		}))
	}

	return logger
}
