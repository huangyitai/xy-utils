package skywalking

import (
	"context"
	"os"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/huangyitai/xy-utils/confx"
	"github.com/huangyitai/xy-utils/xxx"
)

type skywalkingAPI struct {
	span go2sky.Span
	id   string
}

// Tag ...
func (a *skywalkingAPI) Tag(k string, v string) {
	a.span.Tag(go2sky.Tag(k), v)
}

// Log ...
func (a *skywalkingAPI) Log(s ...string) {
	a.span.Log(time.Now(), s...)
}

// ID ...
func (a *skywalkingAPI) ID() string {
	return a.id
}

// NewAPIWithContext 向下兼容，现在不进行任何操作
func NewAPIWithContext(ctx context.Context) context.Context {
	return ctx
}

// SetupTracerFromStringWithPath 从配置字符串初始化tracer
func SetupTracerFromStringWithPath(str, format, path string) (*go2sky.Tracer, error) {
	cfg := NewConfig()
	if path == "" {
		err := confx.UnmarshalAny([]byte(os.ExpandEnv(str)), cfg, format)
		if err != nil {
			return nil, err
		}
	} else {
		err := confx.UnmarshalAnyWithPath([]byte(os.ExpandEnv(str)), cfg, format, path)
		if err != nil {
			return nil, err
		}
	}
	return SetupTracerFromConfig(cfg)
}

// SetupTracerReadConfigWithPath 读取配置初始化tracer
func SetupTracerReadConfigWithPath(read confx.ReadFunc, key, format, path string) (*go2sky.Tracer, error) {
	bs, err := read(key)
	if err != nil {
		return nil, err
	}
	return SetupTracerFromStringWithPath(xxx.UnsafeToString(bs), format, path)
}
