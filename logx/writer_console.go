package logx

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	defaultConsoleIsPretty        = false
	defaultConsoleWriteSync       = false
	defaultConsoleQueueSize       = 10000
	defaultConsoleWriteIntervalMS = 0
	defaultConsoleTimeFormat      = "2006-01-02 15:04:05.000"
)

var stdoutWrapper = WrapWriterWithCloserSync(os.Stdout)

// ConsoleLogWriterConfig 智研日志sdk上报配置
type ConsoleLogWriterConfig struct {
	BaseWriterConfig `yaml:",inline" mapstructure:",squash"`

	Pretty bool `yaml:"pretty" mapstructure:"pretty"`

	TimeFormat string `yaml:"time_format" mapstructure:"time_format"`

	WriteSync bool `yaml:"write_sync" mapstructure:"write_sync"`

	QueueSize       int `yaml:"queue_size" mapstructure:"queue_size"`
	WriteIntervalMS int `yaml:"write_interval_ms" mapstructure:"write_interval_ms"`
}

func fixConsoleLogWriterConfig(cfg *ConsoleLogWriterConfig) {
	fixBaseWriterConfig(&cfg.BaseWriterConfig)
	if cfg.QueueSize <= 0 {
		cfg.QueueSize = defaultConsoleQueueSize
	}
	if cfg.TimeFormat == "" {
		cfg.TimeFormat = defaultConsoleTimeFormat
	}
}

// Setup ...
func (cfg *ConsoleLogWriterConfig) Setup() (zerolog.LevelWriter, error) {
	fixConsoleLogWriterConfig(cfg)

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	var wc io.WriteCloser
	if cfg.Pretty {
		w := zerolog.ConsoleWriter{
			Out:             stdoutWrapper,
			NoColor:         true,
			TimeFormat:      cfg.TimeFormat,
			FormatTimestamp: consoleFormatTimestamp(cfg.TimeFormat),
			FormatLevel:     consoleDefaultFormatLevel(),
		}
		wc = ConsoleWriter{
			ConsoleWriter: w,
			Closer:        stdoutWrapper,
		}
	} else {
		wc = stdoutWrapper
	}

	if !cfg.WriteSync {
		wc = diode.NewWriter(wc, cfg.QueueSize, time.Duration(cfg.WriteIntervalMS)*time.Millisecond, func(missed int) {
			_, _ = fmt.Fprintf(os.Stderr, "Logger dropped %d messages.\n", missed)
		})
	}

	return WrapWithLevel(wc, level), nil
}

// SyncWriteCloserWrapper ...
type SyncWriteCloserWrapper struct {
	file *os.File
	m    sync.Mutex
}

// Write ...
func (w *SyncWriteCloserWrapper) Write(p []byte) (int, error) {
	w.m.Lock()
	defer w.m.Unlock()
	_, err := w.file.Write(p)
	if err != nil {
		fmt.Println("sync write fail ", err)
	}
	return len(p), nil
}

// Close ...
func (w *SyncWriteCloserWrapper) Close() error {
	w.m.Lock()
	defer w.m.Unlock()

	if w.file != nil {
		return w.file.Sync()
	}
	return nil
}

// ConsoleWriter ...
type ConsoleWriter struct {
	zerolog.ConsoleWriter
	io.Closer
}

// WrapWriterWithCloserSync ...
func WrapWriterWithCloserSync(file *os.File) *SyncWriteCloserWrapper {
	return &SyncWriteCloserWrapper{file: file}
}

func consoleFormatTimestamp(timeFormat string) zerolog.Formatter {
	if timeFormat == "" {
		timeFormat = defaultConsoleTimeFormat
	}
	return func(i interface{}) string {
		t := "<nil>"
		switch tt := i.(type) {
		case string:
			ts, err := time.Parse(zerolog.TimeFieldFormat, tt)
			if err != nil {
				t = tt
			} else {
				t = ts.Format(timeFormat)
			}
		case json.Number:
			i, err := tt.Int64()
			if err != nil {
				t = tt.String()
			} else {
				var sec, nsec int64 = i, 0
				switch zerolog.TimeFieldFormat {
				case zerolog.TimeFormatUnixMs:
					nsec = int64(time.Duration(i) * time.Millisecond)
					sec = 0
				case zerolog.TimeFormatUnixMicro:
					nsec = int64(time.Duration(i) * time.Microsecond)
					sec = 0
				}
				ts := time.Unix(sec, nsec).Local()
				t = ts.Format(timeFormat)
			}
		}
		return t
	}
}

func consoleDefaultFormatLevel() zerolog.Formatter {
	return func(i interface{}) string {
		if ll, ok := i.(string); ok {
			switch ll {
			case "trace":
				return "TRACE"
			case "debug":
				return "DEBUG"
			case "info":
				return " INFO"
			case "warn":
				return " WARN"
			case "error":
				return "ERROR"
			case "fatal":
				return "FATAL"
			case "panic":
				return "PANIC"
			default:
				return "?????"
			}
		} else {
			if i == nil {
				return "?????"
			} else {
				return strings.ToUpper(fmt.Sprintf("%s", i))[0:5]
			}
		}
	}
}
