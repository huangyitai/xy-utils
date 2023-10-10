package confx

import (
	"context"
	"reflect"

	"github.com/huangyitai/xy-utils/contx"
	"github.com/huangyitai/xy-utils/deferx"
	"github.com/huangyitai/xy-utils/tagx"
	"github.com/huangyitai/xy-utils/xxx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default ...
var Default = ReadFile

// DefaultFormat ...
var DefaultFormat = JSON

// DefaultWatch ...
var DefaultWatch = func(key string, onUpdated func([]byte) error) error {
	val, err := Default(key)
	if err != nil {
		return err
	}
	return onUpdated(val)
}

// Unmarshal ...
func Unmarshal(key string, ptr interface{}, format string, opts ...viper.DecoderConfigOption) error {
	return Default.Unmarshal(key, ptr, format, opts...)
}

// UnmarshalJSON ...
func UnmarshalJSON(key string, ptr interface{}, opts ...viper.DecoderConfigOption) error {
	return Default.UnmarshalJSON(key, ptr, opts...)
}

// UnmarshalYAML ...
func UnmarshalYAML(key string, ptr interface{}, opts ...viper.DecoderConfigOption) error {
	return Default.UnmarshalYAML(key, ptr, opts...)
}

// UnmarshalTOML ...
func UnmarshalTOML(key string, ptr interface{}, opts ...viper.DecoderConfigOption) error {
	return Default.UnmarshalTOML(key, ptr, opts...)
}

// Configure ...
func Configure(cfg interface{}, opts ...viper.DecoderConfigOption) error {
	return ReadConfig(cfg, opts...)
}

// ConfigureWithContext ...
func ConfigureWithContext(ctx context.Context, cfg interface{}, opts ...viper.DecoderConfigOption) error {
	return ReadConfigWithContext(ctx, cfg, opts...)
}

// ReadConfigWithContext ...
func ReadConfigWithContext(ctx context.Context, cfg interface{}, opts ...viper.DecoderConfigOption) (e error) {
	return contx.CallWithContext(ctx, func() error {
		return ReadConfig(cfg, opts...)
	})
}

// ReadConfig 读入一个config结构体的配置，config中每一个字段都会进行读入
func ReadConfig(cfg interface{}, opts ...viper.DecoderConfigOption) (e error) {
	defer deferx.PanicToError(&e)

	//反射获取结构体指针的值和类型
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	//枚举结构体成员
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)

		err := ReadField(fieldType.Name, fieldValue, fieldType.Tag, opts...)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadField 读入一个字段的配置
func ReadField(name string, value reflect.Value, tag reflect.StructTag, opts ...viper.DecoderConfigOption) error {
	sign := xxx.NewSignStr().WithPath("confx", "ReadField").WithProp("name", name)
	info := TagInfo{Name: name, Format: DefaultFormat}
	err := tagx.UnmarshalTag(TagKey, &info, tag)
	if err != nil {
		return err
	}
	if info.Ignored {
		return nil
	}
	if info.Name == "" {
		info.Name = name
	}
	sign = sign.WithProp(TagKey, info.Name)

	var read ReadFunc
	if info.Read == "" {
		if tb, ok := value.Interface().(Binding); ok {
			read = tb.ReadFunc()
		}
		log.Trace().Str("sName", info.Name).Msgf("%s use Binding.ReadFunc()", sign)
	} else {
		log.Trace().
			Str("sName", info.Name).
			Str("sRead", info.Read).
			Msgf("%s use readFuncTable['%s']", sign, info.Read)

		read = readFuncTable[info.Read]
	}

	if read == nil {
		read = Default
		log.Warn().Str("sName", info.Name).Str("sRead", info.Read).
			Msgf("%s readFunc is nil, use Default", sign)
	}

	if info.Path != "" {
		err = read.ReadValueWithPath(info.Name, value, info.Format, info.Path, opts...)
	} else {
		err = read.ReadValue(info.Name, value, info.Format, opts...)
	}
	log.Err(err).
		Str("sName", info.Name).Str("sFormat", info.Format).
		Msgf("%s ReadField end", sign)

	return err
}

// WatchField ...
func WatchField(name string, watchType reflect.Type, callback func(value reflect.Value) error,
	tag reflect.StructTag, opts ...viper.DecoderConfigOption) error {
	sign := xxx.NewSignStr().WithPath("confx", "WatchField").WithProp("name", name)
	info := TagInfo{Name: name, Format: DefaultFormat}
	err := tagx.UnmarshalTag(TagKey, &info, tag)
	if err != nil {
		return err
	}
	if info.Ignored {
		return nil
	}
	if info.Name == "" {
		info.Name = name
	}
	sign = sign.WithProp(TagKey, info.Name)

	var val reflect.Value
	if watchType.Kind() == reflect.Ptr {
		val = reflect.New(watchType.Elem())
	} else {
		val = reflect.Zero(watchType)
	}

	var watch WatchFunc
	if info.Read == "" {
		if tb, ok := val.Interface().(Binding); ok {
			watch = tb.WatchFunc()
		}
		log.Trace().Str("sName", info.Name).Msgf("%s use Binding.WatchFunc()", sign)
	} else {
		log.Trace().
			Str("sName", info.Name).
			Str("sRead", info.Read).
			Msgf("%s use watchFuncTable['%s']", sign, info.Read)

		watch = watchFuncTable[info.Read]
	}

	if watch == nil {
		watch = DefaultWatch
		log.Warn().Str("sName", info.Name).Str("sRead", info.Read).
			Msgf("%s watchFunc is nil, use DefaultWatch", sign)
	}

	if info.Path != "" {
		err = watch.WatchValueWithPath(info.Name, watchType, callback, info.Format, info.Path, opts...)
	} else {
		err = watch.WatchValue(info.Name, watchType, callback, info.Format, opts...)
	}
	log.Err(err).
		Str("sName", info.Name).Str("sFormat", info.Format).
		Msgf("%s WatchField end", sign)

	return err
}
