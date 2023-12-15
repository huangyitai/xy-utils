package srvx

import (
	"context"
	"fmt"

	"github.com/huangyitai/xy-utils/confx"
	"github.com/huangyitai/xy-utils/contx"
	"github.com/huangyitai/xy-utils/dox"
	"github.com/huangyitai/xy-utils/logx"
	"github.com/huangyitai/xy-utils/metricx"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

// Starter ...
type Starter struct {
	enableSummaryLog   bool
	enableReportMetric bool

	metricName string

	read   confx.ReadFunc
	key    string
	format string
	path   string
}

// NewStarter ...
func NewStarter() *Starter {
	return &Starter{
		enableSummaryLog:   false,
		enableReportMetric: true,
		read:               confx.ReadFile,
		key:                "xy_go.yaml",
		format:             "yaml",
		path:               "xy_go.server",
	}
}

// Start ...
func (s *Starter) Start(ctx context.Context, r contx.ContextRunner) error {
	var err error
	if s.read != nil {
		err = SetupMergeReadWithPath(s.read, s.key, s.format, s.path)
		if err != nil {
			return err
		}
	}

	err = s.checkGlobalConfig()
	if err != nil {
		return err
	}

	err = s.initCron()
	if err != nil {
		return err
	}

	log.Info().Stringer("sConfig", logx.JSONStr(global)).
		Msgf("[srvx/Starter] Service=%s Instance=%s Platform=%s Project=%s Environment=%s Binary=%s starting...",
			global.Service, global.Instance, global.Platform, global.Project, global.Environment, global.Binary)

	err = r(ctx)
	// 多次wait只有首次会阻塞，后续操作会直接返回, 如果后续启动器返回错误，直接中断等待
	e := dox.InterruptOrWaitForCloseSignal(err != nil)
	if e != nil {
		log.Err(e).Msg("[srvx/Starter] dox.WaitForCloseSignal error")
	}
	return err
}

func (s *Starter) initCron() error {
	crn := cron.New(cron.WithSeconds())
	_, err := crn.AddFunc("10/10 * * * * *", s.job10)
	if err != nil {
		log.Err(err).Caller().Msg("[srvx/Starter] cron.AddFunc error")
		return err
	}
	crn.Start()
	dox.SyncDoBeforeClose(func() {
		log.Trace().Msg("[srvx/Starter] cron stopping...")
		crn.Stop()
		log.Trace().Msg("[srvx/Starter] cron stopped")
	})
	return nil
}

func (s *Starter) checkGlobalConfig() error {
	if global.Environment == "" {
		return fmt.Errorf("[srvx/Starter] Environment is empty")
	}
	return nil
}

// EnableSummaryLog ...
func (s *Starter) EnableSummaryLog(enable bool) *Starter {
	s.enableSummaryLog = enable
	return s
}

// EnableReportMetric ...
func (s *Starter) EnableReportMetric(enable bool) *Starter {
	s.enableReportMetric = enable
	return s
}

func (s *Starter) getMetricName() string {
	if s.metricName != "" {
		return s.metricName
	}
	if metricx.APIs[DefaultMetricName] != nil {
		return DefaultMetricName
	}
	if IsDev() {
		return SystemDevMetricName
	}
	if IsPre() {
		return SystemPreMetricName
	}
	if IsProd() {
		return SystemProdMetricName
	}
	return DefaultMetricName
}

// ReadConfigWithPath TODO
func (s *Starter) ReadConfigWithPath(read func(string) ([]byte, error), key, format, path string) *Starter {
	s.read = read
	s.key = key
	s.format = format
	s.path = path
	return s
}

// SetMetricName 设置metric上报配置名
func (s *Starter) SetMetricName(name string) *Starter {
	s.metricName = name
	return s
}
