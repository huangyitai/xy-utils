package confx

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

// ReadValue 从字节数组中读取value的内容，如果value是默认值且支持设置默认值，则会进行默认值设置
func ReadValue(bs []byte, value reflect.Value, format string, opts ...viper.DecoderConfigOption) error {
	value = reflect.Indirect(value)

	// 如果value的指针类型可以初始化默认值（必须要求为指针类型实现NeedInitDefault接口）且value的值为0值，则执行初始化默认值
	if i, ok := value.Addr().Interface().(NeedInitDefault); ok && value.IsZero() {
		i.InitDefault()
	}

	switch value.Interface().(type) {
	case []byte:
		value.Set(reflect.ValueOf(bs))

	case string: //string 类型配置读取
		value.SetString(string(bs))

	case int, int8, int16, int32, int64: //int 读取
		i64, err := cast.ToInt64E(string(bs))
		if err != nil {
			return errors.WithStack(err)
		}
		value.SetInt(i64)

	case bool: //bool 读取
		b, err := cast.ToBoolE(string(bs))
		if err != nil {
			return errors.WithStack(err)
		}
		value.SetBool(b)

	default: //其他类型通过反序列化读取
		//var err error
		//if value.Kind() == reflect.Ptr {
		//	err = UnmarshalAny(bs, value.Interface(), format)
		//} else {
		//	err = UnmarshalAny(bs, value.Addr().Interface(), format)
		//}
		err := UnmarshalAny(bs, value.Addr().Interface(), format, opts...)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// 如果value的指针类型可以进行默认值补充（必须要求为指针类型实现NeedSetDefault接口）则执行默认值补充
	if i, ok := value.Addr().Interface().(NeedSetDefault); ok {
		i.SetDefault()
	}

	return nil
}

// UnmarshalAny ...
func UnmarshalAny(bs []byte, ptr interface{}, format string, opts ...viper.DecoderConfigOption) error {
	format = strings.ToLower(format)
	v := viper.New()
	v.SetConfigType(format)
	err := v.ReadConfig(bytes.NewBuffer(bs))
	if err != nil {
		return err
	}

	return v.Unmarshal(ptr, append(defaultDecoderConfigOptions(), opts...)...)
}

// ReadValueWithPath 从字节数组的指定路径中读取value的内容，如果value是默认值且支持设置默认值，则会进行默认值设置
func ReadValueWithPath(bs []byte, value reflect.Value, format string, path string,
	opts ...viper.DecoderConfigOption) error {

	v := viper.New()
	v.SetConfigType(format)
	err := v.ReadConfig(bytes.NewBuffer(bs))
	if err != nil {
		return err
	}

	value = reflect.Indirect(value)

	// 如果value的指针类型可以初始化默认值（必须要求为指针类型实现NeedInitDefault接口）且value的值为0值，则执行初始化默认值
	if i, ok := value.Addr().Interface().(NeedInitDefault); ok && value.IsZero() {
		i.InitDefault()
	}

	switch value.Interface().(type) {
	case []byte:
		value.Set(reflect.ValueOf([]byte(v.GetString(path))))

	case string: //string 类型配置读取
		value.SetString(v.GetString(path))

	case int, int8, int16, int32, int64: //int 读取
		i64, err := cast.ToInt64E(v.GetInt64(path))
		if err != nil {
			return errors.WithStack(err)
		}
		value.SetInt(i64)

	case bool: //bool 读取
		b, err := cast.ToBoolE(v.GetBool(path))
		if err != nil {
			return errors.WithStack(err)
		}
		value.SetBool(b)

	default: //其他类型通过反序列化读取
		//if value.Kind() == reflect.Ptr {
		//	err = v.UnmarshalKey(path, value.Interface())
		//} else {
		//	err = v.UnmarshalKey(path, value.Addr().Interface())
		//}
		err = v.UnmarshalKey(path, value.Addr().Interface(), append(defaultDecoderConfigOptions(), opts...)...)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// 如果value的指针类型可以进行默认值补充（必须要求为指针类型实现NeedSetDefault接口）则执行默认值补充
	if i, ok := value.Addr().Interface().(NeedSetDefault); ok {
		i.SetDefault()
	}

	return nil
}

// UnmarshalAnyWithPath ...
func UnmarshalAnyWithPath(bs []byte, ptr interface{}, format, path string,
	opts ...viper.DecoderConfigOption) error {

	format = strings.ToLower(format)
	v := viper.New()
	v.SetConfigType(format)
	err := v.ReadConfig(bytes.NewBuffer(bs))
	if err != nil {
		return err
	}

	return v.UnmarshalKey(path, ptr, append(defaultDecoderConfigOptions(), opts...)...)
}
