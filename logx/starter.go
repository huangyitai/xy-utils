package logx

import (
	"context"

	"github.com/huangyitai/xy-utils/confx"
	"github.com/huangyitai/xy-utils/contx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Starter ...
type Starter struct {
	skipConfig bool
	hooks      []zerolog.Hook

	read   confx.ReadFunc
	key    string
	format string
	path   string
}

// NewStarter ...
func NewStarter() *Starter {
	return &Starter{
		skipConfig: false,
		read:       confx.ReadFile,
		key:        "xy_log.yaml",
		format:     "yaml",
	}
}

// Start ...
func (s *Starter) Start(ctx context.Context, r contx.ContextRunner) error {
	var err error
	if !s.skipConfig && s.read != nil {
		if s.path == "" {
			err = SetupReadConfig(s.read, s.key, s.format)
		} else {
			err = SetupReadConfigWithPath(s.read, s.key, s.format, s.path)
		}
		if err != nil {
			return err
		}

		// 添加全局日志钩子
		for _, hook := range s.hooks {
			log.Logger = log.Logger.Hook(hook)
		}
	}
	err = r(ctx)
	CloseAndWait()
	return err
}

// ReadConfig ...
func (s *Starter) ReadConfig(read func(string) ([]byte, error), key, format string) *Starter {
	s.read = read
	s.key = key
	s.format = format
	return s
}

// ReadConfigWithPath ...
func (s *Starter) ReadConfigWithPath(read func(string) ([]byte, error), key, format, path string) *Starter {
	s.read = read
	s.key = key
	s.format = format
	s.path = path
	return s
}

// ReadFrom ...
func (s *Starter) ReadFrom(read func(string) ([]byte, error)) *Starter {
	s.read = read
	return s
}

// ReadKey ...
func (s *Starter) ReadKey(key string) *Starter {
	s.key = key
	return s
}

// ReadPath ...
func (s *Starter) ReadPath(path string) *Starter {
	s.path = path
	return s
}

// SkipConfig ...
func (s *Starter) SkipConfig(skip bool) *Starter {
	s.skipConfig = skip
	return s
}

// AddGlobalHooks 添加全局日志钩子函数
func (s *Starter) AddGlobalHooks(hooks ...zerolog.Hook) *Starter {
	s.hooks = append(s.hooks, hooks...)
	return s
}
