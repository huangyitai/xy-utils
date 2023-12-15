package logx

import (
	"io"
	"path/filepath"

	"github.com/huangyitai/xy-utils/logx/thrid_party/rollwriter"
	"github.com/rs/zerolog"
)

// 时间单位配置字段
const (
	fileTimeUnitMinute = "minute"
	fileTimeUnitHour   = "hour"
	fileTimeUnitDay    = "day"
	fileTimeUnitMonth  = "month"
	fileTimeUnitYear   = "year"

	fileWriteModeSync  = "sync"
	fileWriteModeAsync = "async"
	fileWriteModeFast  = "fast"

	fileRollTypeSize = "size"
	fileRollTypeTime = "time"

	// TimeFormatMinute 分钟
	fileTimeFormatMinute = "%Y%m%d%H%M"
	// TimeFormatHour 小时
	fileTimeFormatHour = "%Y%m%d%H"
	// TimeFormatDay 天
	fileTimeFormatDay = "%Y%m%d"
	// TimeFormatMonth 月
	fileTimeFormatMonth = "%Y%m"
	// TimeFormatYear 年
	fileTimeFormatYear = "%Y"

	defaultFileLogPath         = "/usr/local/xy/log"
	defaultFileFileName        = "xy.log"
	defaultFileRollType        = fileRollTypeTime
	defaultFileWriteMode       = fileWriteModeAsync
	defaultFileMaxAge          = 0
	defaultFileMaxBackups      = 0
	defaultFileMaxSize         = 100
	defaultFileTimeUnit        = fileTimeUnitDay
	defaultFileQueueSize       = 10000
	defaultFileBufferSize      = 4 * 1024
	defaultFileWriteIntervalMS = 100
)

// TimeUnit 文件按时间分割的时间单位，支持：minute/hour/day/month/year
type TimeUnit string

// Format 返回时间单位的格式字符串（c风格），默认返回day的格式字符串
func (t TimeUnit) Format() string {
	var timeFmt string
	switch t {
	case fileTimeUnitMinute:
		timeFmt = fileTimeFormatMinute
	case fileTimeUnitHour:
		timeFmt = fileTimeFormatHour
	case fileTimeUnitDay:
		timeFmt = fileTimeFormatDay
	case fileTimeUnitMonth:
		timeFmt = fileTimeFormatMonth
	case fileTimeUnitYear:
		timeFmt = fileTimeFormatYear
	default:
		timeFmt = fileTimeFormatDay
	}
	return "." + timeFmt
}

// FileLogWriterConfig 本地文件的配置
type FileLogWriterConfig struct {
	BaseWriterConfig `yaml:",inline" mapstructure:",squash"`

	// LogPath 日志路径名  /usr/local/trpc/log/
	LogPath string `yaml:"log_path" mapstructure:"log_path"`
	// Filename 日志路径文件名  trpc.log
	Filename string `yaml:"filename" mapstructure:"filename"`
	// RollType 文件滚动类型，size-按大小分割文件，time-按时间分割文件，默认按大小分割
	RollType string `yaml:"roll_type" mapstructure:"roll_type"`
	// WriteMode 日志写入模式，1-同步，2-异步，3-极速(异步丢弃)
	WriteMode string `yaml:"write_mode" mapstructure:"write_mode"`
	// MaxAge 日志最大保留时间, 天
	MaxAge int `yaml:"max_age" mapstructure:"max_age"`
	// MaxBackups 日志最大文件数
	MaxBackups int `yaml:"max_backups" mapstructure:"max_backups"`
	// Compress 日志文件是否压缩
	Compress bool `yaml:"compress" mapstructure:"compress"`

	// 以下参数按大小分割时才有效
	// MaxSize 日志文件最大大小（单位MB）
	MaxSize int `yaml:"max_size_mb" mapstructure:"max_size_mb"`

	QueueSize       int `yaml:"queue_size" mapstructure:"queue_size"`
	BufferSize      int `yaml:"buffer_size" mapstructure:"buffer_size"`
	WriteIntervalMS int `yaml:"write_interval_ms" mapstructure:"write_interval_ms"`

	// 以下参数按时间分割时才有效
	// TimeUnit 按时间分割文件的时间单位
	// 支持year/month/day/hour/minute, 默认为day
	TimeUnit TimeUnit `yaml:"time_unit" mapstructure:"time_unit"`
}

func fixFileLogWriterConfig(cfg *FileLogWriterConfig) {
	fixBaseWriterConfig(&cfg.BaseWriterConfig)

	if cfg.LogPath == "" {
		cfg.LogPath = defaultFileLogPath
	}
	if cfg.Filename == "" {
		cfg.Filename = defaultFileFileName
	}
	if cfg.RollType == "" {
		cfg.RollType = defaultFileRollType
	}
	if cfg.WriteMode == "" {
		cfg.WriteMode = defaultFileWriteMode
	}
	if cfg.MaxAge < 0 {
		cfg.MaxAge = defaultFileMaxAge
	}
	if cfg.MaxBackups < 0 {
		cfg.MaxBackups = defaultFileMaxBackups
	}
	if cfg.MaxSize <= 0 {
		cfg.MaxSize = defaultFileMaxSize
	}
	if cfg.TimeUnit == "" {
		cfg.TimeUnit = defaultFileTimeUnit
	}
	if cfg.QueueSize <= 0 {
		cfg.QueueSize = defaultFileQueueSize
	}
	if cfg.BufferSize <= 0 {
		cfg.BufferSize = defaultFileBufferSize
	}
	if cfg.WriteIntervalMS <= 0 {
		cfg.WriteIntervalMS = defaultFileWriteIntervalMS
	}
}

// Setup ...
func (cfg *FileLogWriterConfig) Setup() (zerolog.LevelWriter, error) {
	fixFileLogWriterConfig(cfg)

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	var rw io.WriteCloser

	filename := cfg.Filename
	if cfg.LogPath != "" {
		filename = filepath.Join(cfg.LogPath, filename)
	}

	if cfg.RollType == fileRollTypeSize {
		rw, err = rollwriter.NewRollWriter(
			filename,
			rollwriter.WithMaxAge(cfg.MaxAge),
			rollwriter.WithMaxBackups(cfg.MaxBackups),
			rollwriter.WithCompress(cfg.Compress),
			rollwriter.WithMaxSize(cfg.MaxSize),
		)
		if err != nil {
			return nil, err
		}
	} else {
		rw, err = rollwriter.NewRollWriter(
			filename,
			rollwriter.WithMaxAge(cfg.MaxAge),
			rollwriter.WithMaxBackups(cfg.MaxBackups),
			rollwriter.WithCompress(cfg.Compress),
			rollwriter.WithMaxSize(cfg.MaxSize),
			rollwriter.WithRotationTime(cfg.TimeUnit.Format()),
		)
		if err != nil {
			return nil, err
		}
	}

	if cfg.WriteMode != fileWriteModeSync {
		dropLog := cfg.WriteMode == fileWriteModeFast
		rw = rollwriter.NewAsyncRollWriter(rw, rollwriter.WithDropLog(dropLog))
	}

	return WrapWithLevel(rw, level), nil
}
