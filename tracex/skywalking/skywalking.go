package skywalking

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/huangyitai/xy-utils/tracex"
)

// SpanContextKey 存放span的context key（保持兼容）
const SpanContextKey = "sky-span"

// DefaultComponentId 默认图标为grpc
const DefaultComponentId = 23

// Tracer 全局tracer对象
var Tracer *go2sky.Tracer

// GetSpan return span instance
func GetSpan(ctx context.Context) (go2sky.Span, error) {
	span, ok := ctx.Value(SpanContextKey).(go2sky.Span)
	if !ok {
		return nil, errors.New("get span from context failed")
	}

	return span, nil
}

// WithCtx 将span存入context，并初始化tracex.API
func WithCtx(ctx context.Context, span go2sky.Span) context.Context {
	ctx = context.WithValue(ctx, SpanContextKey, span)
	api := &skywalkingAPI{
		span: span,
		id:   go2sky.TraceID(ctx),
	}
	return tracex.WithCtx(ctx, api)
}

// SetupTracerFromConfig 从Config对象初始化Tracer
func SetupTracerFromConfig(cfg *Config) (*go2sky.Tracer, error) {
	opts, err := generateTracerOptions(cfg)
	if err != nil {
		return nil, err
	}

	skyTracer, err := go2sky.NewTracer(cfg.Service, opts...)
	if err != nil {
		return nil, err
	}
	return skyTracer, nil
}

// generateReporterOptions 自定义配置转换为skyWalking配置
func generateReporterOptions(config *Config) (opts []reporter.GRPCReporterOption, err error) {
	result := make([]reporter.GRPCReporterOption, 0, 4)
	if config.CheckInterval != "" {
		t, err := time.ParseDuration(config.CheckInterval)
		if err != nil {
			return result, err
		}
		result = append(result, reporter.WithCheckInterval(t))
	}

	result = append(result, reporter.WithInstanceProps(config.InstanceProps))
	if config.MaxSendQueueSize != 0 {
		result = append(result, reporter.WithMaxSendQueueSize(config.MaxSendQueueSize))
	}
	if config.Auth != "" {
		result = append(result, reporter.WithAuthentication(config.Auth))
	}
	if config.ComponentId == 0 {
		config.ComponentId = DefaultComponentId
	}
	return result, nil
}

func generateTracerOptions(cfg *Config) ([]go2sky.TracerOption, error) {
	opts, err := generateReporterOptions(cfg)
	if err != nil {
		return nil, err
	}

	report, err := reporter.NewGRPCReporter(cfg.Address, opts...)
	if err != nil {
		return nil, err
	}

	ins := os.Getenv("HOSTNAME")
	if len(ins) == 0 {
		ins = "unknow-instance"
	}

	res := []go2sky.TracerOption{
		go2sky.WithReporter(report), go2sky.WithInstance(ins), go2sky.WithSampler(cfg.SamplingRate),
	}
	return res, nil
}
