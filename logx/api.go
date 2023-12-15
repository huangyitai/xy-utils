package logx

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/huangyitai/xy-utils/confx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// WithSubLogger 创建一个固定包含指定内容的子context logger
func WithSubLogger(ctx context.Context, f func(c zerolog.Context) zerolog.Context) context.Context {
	logger := log.Ctx(ctx)
	c := logger.With()
	if f != nil {
		c = f(c)
	}
	sub := c.Logger()
	return sub.WithContext(ctx)
}

// WithSubLevel 创建一个最低输出级别为指定级别的子context logger
func WithSubLevel(ctx context.Context, lv zerolog.Level) context.Context {
	sub := log.Ctx(ctx).Level(lv)
	return sub.WithContext(ctx)
}

// SetupReadConfigWithPath ...
func SetupReadConfigWithPath(read confx.ReadFunc, key, format, path string) error {
	bs, err := read(key)
	if err != nil {
		return err
	}
	return SetupFromStringWithPath(string(bs), format, path)
}

// SetupReadConfig ...
func SetupReadConfig(read confx.ReadFunc, key, format string) error {
	bs, err := read(key)
	if err != nil {
		return err
	}
	return SetupFromString(string(bs), format)
}

// SetupFromString ...
func SetupFromString(config string, format string) error {
	cfg := new(Config)
	err := confx.UnmarshalAny([]byte(os.ExpandEnv(config)), cfg, format)
	if err != nil {
		return err
	}
	return SetupFromConfig(cfg)
}

// SetupFromStringWithPath ...
func SetupFromStringWithPath(config string, format, path string) error {
	cfg := new(Config)
	err := confx.UnmarshalAnyWithPath([]byte(os.ExpandEnv(config)), cfg, format, path)
	if err != nil {
		return err
	}
	return SetupFromConfig(cfg)
}

// SetupFromConfig ...
func SetupFromConfig(cfg *Config) error {
	Close()

	lock.Lock()
	defer lock.Unlock()

	writers, err := setupWriters(cfg.WriterConfigs)
	if err != nil {
		return err
	}

	setupFormat(&cfg.FormatConfig)

	mw := zerolog.MultiLevelWriter(writers...)
	log.Logger = setupGlobal(&cfg.GlobalConfig, mw)
	zerolog.DefaultContextLogger = &log.Logger

	// 设置全局日志级别，不配置默认为trace
	if cfg.GlobalConfig.Level != "" {
		gLevel, err := zerolog.ParseLevel(cfg.GlobalConfig.Level)
		if err != nil {
			return err
		}
		zerolog.SetGlobalLevel(gLevel)
		log.Info().Msgf("[logx/SetupFromConfig] set GlobalLevel = %s", cfg.GlobalConfig.Level)
	} else {
		log.Info().Msg("[logx/SetupFromConfig] default GlobalLevel = trace")
	}

	gWriters = writers

	return nil
}

// Close ...
func Close() {
	lock.Lock()
	defer lock.Unlock()

	if len(gWriters) != 0 {
		for _, writer := range gWriters {
			if closer, ok := writer.(io.Closer); ok {
				_ = closer.Close()
			}
		}
		gWriters = nil
		l := zerolog.Nop()
		zerolog.DefaultContextLogger = &l
		log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}
}

// CloseAndWait ...
func CloseAndWait() {
	if gWriters != nil {
		log.Trace().Msg("[logx] closing...")
		time.Sleep(2 * time.Second)
		Close()
		//XXX 防止未完成输出的日志丢失，等待固定时间2s，未来应该采用更加优雅的方式
		time.Sleep(1 * time.Second)
	}
}
