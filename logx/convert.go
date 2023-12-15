package logx

import (
	"fmt"
	"net"
	"time"

	"github.com/huangyitai/xy-utils/xxx"
	"github.com/iancoleman/strcase"
	"github.com/rs/zerolog"
)

// KeyPrefixString ...
var (
	KeyPrefixString  = "s"
	KeyPrefixBoolean = "b"
	KeyPrefixInteger = "i"
	KeyPrefixFloat   = "f"
)

// Key ...
func Key(prefix, key string) string {
	return prefix + strcase.ToCamel(key)
}

// Any ...
func Any(key string, value interface{}) *AnyValue {
	return &AnyValue{
		k: key,
		v: value,
	}
}

// AnyValue ...
type AnyValue struct {
	k string
	v interface{}
}

// Prefix ...
func (a *AnyValue) Prefix() string {
	switch a.v.(type) {
	case string, []byte, error, []error, *string,
		[]string, []bool, []int, []int8, []int16, []int32, []int64, []uint, []uint16, []uint32, []uint64,
		[]float32, []float64, []time.Time, []time.Duration, nil, net.IP, net.IPNet, net.HardwareAddr, fmt.Stringer:
		return KeyPrefixString
	case bool, *bool:
		return KeyPrefixBoolean
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
		*int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64:
		return KeyPrefixInteger
	case float32, float64, *float32, *float64:
		return KeyPrefixFloat
	case time.Time, *time.Time:
		switch zerolog.TimeFieldFormat {
		case zerolog.TimeFormatUnix, zerolog.TimeFormatUnixMs, zerolog.TimeFormatUnixMicro:
			return KeyPrefixInteger
		default:
			return KeyPrefixString
		}
	case time.Duration, *time.Duration:
		if zerolog.DurationFieldInteger {
			return KeyPrefixInteger
		} else {
			return KeyPrefixFloat
		}
	default:
		return KeyPrefixString
	}
}

// Key ...
func (a *AnyValue) Key() string {
	return Key(a.Prefix(), a.k)
}

// MarshalZerologObject ...
func (a *AnyValue) MarshalZerologObject(e *zerolog.Event) {
	// 首先对空值进行统一处理
	if a.v == nil {
		e.Interface(a.Key(), nil)
		return
	}
	// 类型打表处理
	switch val := a.v.(type) {
	case string:
		e.Str(a.Key(), val)
	case []byte:
		e.Str(a.Key(), string(val))
	case error:
		e.AnErr(a.Key(), val)
	case []error:
		e.Str(a.Key(), xxx.ErrorsToJSONStr(val))
	case bool:
		e.Bool(a.Key(), val)
	case int:
		e.Int(a.Key(), val)
	case int8:
		e.Int8(a.Key(), val)
	case int16:
		e.Int16(a.Key(), val)
	case int32:
		e.Int32(a.Key(), val)
	case int64:
		e.Int64(a.Key(), val)
	case uint:
		e.Uint(a.Key(), val)
	case uint8:
		e.Uint8(a.Key(), val)
	case uint16:
		e.Uint16(a.Key(), val)
	case uint32:
		e.Uint32(a.Key(), val)
	case uint64:
		e.Uint64(a.Key(), val)
	case float32:
		e.Float32(a.Key(), val)
	case float64:
		e.Float64(a.Key(), val)
	case time.Time:
		e.Time(a.Key(), val)
	case time.Duration:
		e.Dur(a.Key(), val)
	case *string:
		e.Str(a.Key(), *val)
	case *bool:
		e.Bool(a.Key(), *val)
	case *int:
		e.Int(a.Key(), *val)
	case *int8:
		e.Int8(a.Key(), *val)
	case *int16:
		e.Int16(a.Key(), *val)
	case *int32:
		e.Int32(a.Key(), *val)
	case *int64:
		e.Int64(a.Key(), *val)
	case *uint:
		e.Uint(a.Key(), *val)
	case *uint8:
		e.Uint8(a.Key(), *val)
	case *uint16:
		e.Uint16(a.Key(), *val)
	case *uint32:
		e.Uint32(a.Key(), *val)
	case *uint64:
		e.Uint64(a.Key(), *val)
	case *float32:
		e.Float32(a.Key(), *val)
	case *float64:
		e.Float64(a.Key(), *val)
	case *time.Time:
		e.Time(a.Key(), *val)
	case *time.Duration:
		e.Dur(a.Key(), *val)
	case []string, []bool, []int, []int8, []int16, []int32, []int64, []uint, []uint16, []uint32, []uint64,
		[]float32, []float64, []time.Time, []time.Duration:
		e.Stringer(a.Key(), JSONStr(val))
	case net.IP:
		e.IPAddr(a.Key(), val)
	case net.IPNet:
		e.IPPrefix(a.Key(), val)
	case net.HardwareAddr:
		e.MACAddr(a.Key(), val)
	case fmt.Stringer:
		e.Stringer(a.Key(), val)
	default:
		e.Stringer(a.Key(), JSONStr(val))
	}
}
