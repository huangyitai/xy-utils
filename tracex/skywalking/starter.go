package skywalking

import (
	"context"

	"github.com/huangyitai/xy-utils/confx"
	"github.com/huangyitai/xy-utils/contx"
)

// Starter skywalking启动器
type Starter struct {
	skipSetupTracer bool

	read   confx.ReadFunc
	key    string
	format string
	path   string
}

// NewStarter 创建Starter并设置初始值
func NewStarter() *Starter {
	return &Starter{
		skipSetupTracer: false,
		read:            confx.ReadFile,
		key:             "xy_go.yaml",
		format:          "yaml",
		path:            "skywalking",
	}
}

// Start ...
func (s *Starter) Start(ctx context.Context, r contx.ContextRunner) error {
	if !s.skipSetupTracer && s.read != nil {
		tracer, err := SetupTracerReadConfigWithPath(s.read, s.key, s.format, s.path)
		if err != nil {
			return err
		}
		Tracer = tracer
	}
	return r(ctx)
}

// ReadConfigWithPath ...
func (s *Starter) ReadConfigWithPath(read func(string) ([]byte, error), key, format, path string) *Starter {
	s.read = read
	s.key = key
	s.format = format
	s.path = path
	return s
}

// SkipSetupTracer 配置是否跳过初始化Tracer
func (s *Starter) SkipSetupTracer(skip bool) *Starter {
	s.skipSetupTracer = skip
	return s
}
