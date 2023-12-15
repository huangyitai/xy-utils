package logx

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/huangyitai/xy-utils/confx"
	"github.com/huangyitai/xy-utils/xxx"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

const (
	defaultLevel = "debug"

	defaultTimeKey       = "@time"
	defaultTimeFormat    = "2006-01-02 15:04:05.000000"
	defaultLevelKey      = "@level"
	defaultCallerKey     = "@caller"
	defaultStacktraceKey = "@stacktrace"
	defaultMessageKey    = "@message"
	defaultErrorKey      = "@error"
	defaultOffsetKey     = "@offset"

	zhiyanWriterType  = "zhiyan"
	consoleWriterType = "console"
	fileWriterType    = "file"
)

func init() {
	//注册反序列化WriterConfig接口的Hook函数
	confx.AddTypeHook(DecodeWriteConfig)
}

// Config ...
type Config struct {
	GlobalConfig  GlobalConfig   `yaml:"global" mapstructure:"global"`
	FormatConfig  FormatConfig   `yaml:"format" mapstructure:"format"`
	WriterConfigs []WriterConfig `yaml:"writer" mapstructure:"writer"`
}

// DecodeWriteConfig ...
func DecodeWriteConfig(tf reflect.Type, tt reflect.Type, data interface{}) (interface{}, error) {
	//确保是将map写入WriterConfig
	if tf.Kind() != reflect.Map || tt != reflect.TypeOf(new(WriterConfig)).Elem() {
		return data, nil
	}

	v, ok := xxx.MapGet(data, "type")
	if !ok {
		return nil, fmt.Errorf("logx config err: no 'type' present")
	}

	t, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("logx config err: 'type' is not string")
	}

	i, err := newConfigByType(t)
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, i)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// UnmarshalYAML ...
func (c *Config) UnmarshalYAML(value *yaml.Node) error {
	tc := struct {
		GlobalConfig      GlobalConfig `yaml:"global"`
		FormatConfig      FormatConfig `yaml:"format"`
		WriterConfigNodes []yaml.Node  `yaml:"writer"`
	}{}
	err := value.Decode(&tc)
	if err != nil {
		return err
	}

	var writerConfigs []WriterConfig
	for _, node := range tc.WriterConfigNodes {
		bc := BaseWriterConfig{}
		err = node.Decode(&bc)
		if err != nil {
			return err
		}

		i, err := newConfigByType(bc.Type)
		if err != nil {
			return err
		}

		err = node.Decode(i)
		if err != nil {
			return err
		}

		if writerConfig, ok := i.(WriterConfig); ok {
			writerConfigs = append(writerConfigs, writerConfig)
		} else {
			return fmt.Errorf("unexpected writer config type %T", i)
		}
	}

	c.GlobalConfig = tc.GlobalConfig
	c.FormatConfig = tc.FormatConfig
	c.WriterConfigs = writerConfigs
	return nil
}

func newConfigByType(t string) (interface{}, error) {
	switch t {
	case consoleWriterType:
		return &ConsoleLogWriterConfig{}, nil
	case fileWriterType:
		return &FileLogWriterConfig{}, nil
	default:
		return nil, fmt.Errorf("unknown writer type %s", t)
	}
}

// FormatConfig ...
type FormatConfig struct {
	TimeKey       string `yaml:"time_key" mapstructure:"time_key"`
	TimeFormat    string `yaml:"time_field_format" mapstructure:"time_field_format"`
	LevelKey      string `yaml:"level_key" mapstructure:"level_key"`
	CallerKey     string `yaml:"caller_key" mapstructure:"caller_key"`
	StacktraceKey string `yaml:"stacktrace_key" mapstructure:"stacktrace_key"`
	MessageKey    string `yaml:"message_key" mapstructure:"message_key"`
	ErrorKey      string `yaml:"error_key" mapstructure:"error_key"`
	OffsetKey     string `yaml:"offset_key" mapstructure:"offset_key"`
}

func fixFormatConfig(cfg *FormatConfig) {
	if cfg.TimeKey == "" {
		cfg.TimeKey = defaultTimeKey
	}
	if cfg.TimeFormat == "" {
		cfg.TimeFormat = defaultTimeFormat
	}
	if cfg.LevelKey == "" {
		cfg.LevelKey = defaultLevelKey
	}
	if cfg.CallerKey == "" {
		cfg.CallerKey = defaultCallerKey
	}
	if cfg.StacktraceKey == "" {
		cfg.StacktraceKey = defaultStacktraceKey
	}
	if cfg.MessageKey == "" {
		cfg.MessageKey = defaultMessageKey
	}
	if cfg.ErrorKey == "" {
		cfg.ErrorKey = defaultErrorKey
	}
	if cfg.OffsetKey == "" {
		cfg.OffsetKey = defaultOffsetKey
	}
}

// GlobalConfig ...
type GlobalConfig struct {
	Level             string            `yaml:"level" mapstructure:"level"`
	DisableCaller     bool              `yaml:"disable_caller" mapstructure:"disable_caller"`
	DisableErrorStack bool              `yaml:"disable_error_stack" mapstructure:"disable_error_stack"`
	DisableTimestamp  bool              `yaml:"disable_timestamp" mapstructure:"disable_timestamp"`
	DisableOffset     bool              `yaml:"disable_offset" mapstructure:"disable_offset"`
	CustomFields      map[string]string `yaml:"fields" mapstructure:"fields"`
}

// WriterConfig ...
type WriterConfig interface {
	Setup() (zerolog.LevelWriter, error)
}

// BaseWriterConfig ...
type BaseWriterConfig struct {
	Type  string `yaml:"type" mapstructure:"type"`
	Level string `yaml:"level" mapstructure:"level"`
}

func fixBaseWriterConfig(cfg *BaseWriterConfig) {
	if cfg.Level == "" {
		cfg.Level = defaultLevel
	} else {
		cfg.Level = strings.ToLower(cfg.Level)
	}
}
