package logx

import (
	"fmt"
	"github.com/rs/zerolog"
	"net"
	"time"
)

// LogEvent 日志事件
type LogEvent struct {
	event *zerolog.Event
}

// Fields ...
func (e *LogEvent) Fields(fields interface{}) *LogEvent {
	e.event.Fields(fields)
	return e
}

// Dict TODO
func (e *LogEvent) Dict(key string, dict *zerolog.Event) *LogEvent {
	e.event.Dict(key, dict)
	return e
}

// Array TODO
func (e *LogEvent) Array(key string, arr zerolog.LogArrayMarshaler) *LogEvent {
	e.event.Array(key, arr)
	return e
}

// Object TODO
func (e *LogEvent) Object(key string, obj zerolog.LogObjectMarshaler) *LogEvent {
	e.event.Object(key, obj)
	return e
}

// Func TODO
func (e *LogEvent) Func(f func(e *zerolog.Event)) *LogEvent {
	e.event.Func(f)
	return e
}

// EmbedObject TODO
func (e *LogEvent) EmbedObject(obj zerolog.LogObjectMarshaler) *LogEvent {
	e.event.EmbedObject(obj)
	return e
}

// Str TODO
func (e *LogEvent) Str(key, val string) *LogEvent {
	e.event.Str(key, val)
	return e
}

// Strs TODO
func (e *LogEvent) Strs(key string, vals []string) *LogEvent {
	e.event.Strs(key, vals)
	return e
}

// Stringer TODO
func (e *LogEvent) Stringer(key string, val fmt.Stringer) *LogEvent {
	e.event.Stringer(key, val)
	return e
}

// Stringers TODO
func (e *LogEvent) Stringers(key string, vals []fmt.Stringer) *LogEvent {
	e.event.Stringers(key, vals)
	return e
}

// Bytes TODO
func (e *LogEvent) Bytes(key string, val []byte) *LogEvent {
	e.event.Bytes(key, val)
	return e
}

// Hex TODO
func (e *LogEvent) Hex(key string, val []byte) *LogEvent {
	e.event.Hex(key, val)
	return e
}

// RawJSON TODO
func (e *LogEvent) RawJSON(key string, b []byte) *LogEvent {
	e.event.RawJSON(key, b)
	return e
}

// AnErr TODO
func (e *LogEvent) AnErr(key string, err error) *LogEvent {
	e.event.AnErr(key, err)
	return e
}

// Errs TODO
func (e *LogEvent) Errs(key string, errs []error) *LogEvent {
	e.event.Errs(key, errs)
	return e
}

// Err TODO
func (e *LogEvent) Err(err error) *LogEvent {
	e.event.Err(err)
	return e
}

// Stack TODO
func (e *LogEvent) Stack() *LogEvent {
	e.event.Stack()
	return e
}

// Bool TODO
func (e *LogEvent) Bool(key string, b bool) *LogEvent {
	e.event.Bool(key, b)
	return e
}

// Bools TODO
func (e *LogEvent) Bools(key string, b []bool) *LogEvent {
	e.event.Bools(key, b)
	return e
}

// Int TODO
func (e *LogEvent) Int(key string, i int) *LogEvent {
	e.event.Int(key, i)
	return e
}

// Ints TODO
func (e *LogEvent) Ints(key string, i []int) *LogEvent {
	e.event.Ints(key, i)
	return e
}

// Int8 TODO
func (e *LogEvent) Int8(key string, i int8) *LogEvent {
	e.event.Int8(key, i)
	return e
}

// Ints8 TODO
func (e *LogEvent) Ints8(key string, i []int8) *LogEvent {
	e.event.Ints8(key, i)
	return e
}

// Int16 TODO
func (e *LogEvent) Int16(key string, i int16) *LogEvent {
	e.event.Int16(key, i)
	return e
}

// Ints16 TODO
func (e *LogEvent) Ints16(key string, i []int16) *LogEvent {
	e.event.Ints16(key, i)
	return e
}

// Int32 TODO
func (e *LogEvent) Int32(key string, i int32) *LogEvent {
	e.event.Int32(key, i)
	return e
}

// Ints32 TODO
func (e *LogEvent) Ints32(key string, i []int32) *LogEvent {
	e.event.Ints32(key, i)
	return e
}

// Int64 TODO
func (e *LogEvent) Int64(key string, i int64) *LogEvent {
	e.event.Int64(key, i)
	return e
}

// Ints64 TODO
func (e *LogEvent) Ints64(key string, i []int64) *LogEvent {
	e.event.Ints64(key, i)
	return e
}

// Uint TODO
func (e *LogEvent) Uint(key string, i uint) *LogEvent {
	e.event.Uint(key, i)
	return e
}

// Uints TODO
func (e *LogEvent) Uints(key string, i []uint) *LogEvent {
	e.event.Uints(key, i)
	return e
}

// Uint8 TODO
func (e *LogEvent) Uint8(key string, i uint8) *LogEvent {
	e.event.Uint8(key, i)
	return e
}

// Uints8 TODO
func (e *LogEvent) Uints8(key string, i []uint8) *LogEvent {
	e.event.Uints8(key, i)
	return e
}

// Uint16 TODO
func (e *LogEvent) Uint16(key string, i uint16) *LogEvent {
	e.event.Uint16(key, i)
	return e
}

// Uints16 TODO
func (e *LogEvent) Uints16(key string, i []uint16) *LogEvent {
	e.event.Uints16(key, i)
	return e
}

// Uint32 TODO
func (e *LogEvent) Uint32(key string, i uint32) *LogEvent {
	e.event.Uint32(key, i)
	return e
}

// Uints32 TODO
func (e *LogEvent) Uints32(key string, i []uint32) *LogEvent {
	e.event.Uints32(key, i)
	return e
}

// Uint64 TODO
func (e *LogEvent) Uint64(key string, i uint64) *LogEvent {
	e.event.Uint64(key, i)
	return e
}

// Uints64 TODO
func (e *LogEvent) Uints64(key string, i []uint64) *LogEvent {
	e.event.Uints64(key, i)
	return e
}

// Float32 TODO
func (e *LogEvent) Float32(key string, f float32) *LogEvent {
	e.event.Float32(key, f)
	return e
}

// Floats32 TODO
func (e *LogEvent) Floats32(key string, f []float32) *LogEvent {
	e.event.Floats32(key, f)
	return e
}

// Float64 TODO
func (e *LogEvent) Float64(key string, f float64) *LogEvent {
	e.event.Float64(key, f)
	return e
}

// Floats64 TODO
func (e *LogEvent) Floats64(key string, f []float64) *LogEvent {
	e.event.Floats64(key, f)
	return e
}

// Timestamp TODO
func (e *LogEvent) Timestamp() *LogEvent {
	e.event.Timestamp()
	return e
}

// Time TODO
func (e *LogEvent) Time(key string, t time.Time) *LogEvent {
	e.event.Time(key, t)
	return e
}

// Times TODO
func (e *LogEvent) Times(key string, t []time.Time) *LogEvent {
	e.event.Times(key, t)
	return e
}

// Dur TODO
func (e *LogEvent) Dur(key string, d time.Duration) *LogEvent {
	e.event.Dur(key, d)
	return e
}

// Durs TODO
func (e *LogEvent) Durs(key string, d []time.Duration) *LogEvent {
	e.event.Durs(key, d)
	return e
}

// TimeDiff TODO
func (e *LogEvent) TimeDiff(key string, t time.Time, start time.Time) *LogEvent {
	e.event.TimeDiff(key, t, start)
	return e
}

// Interface TODO
func (e *LogEvent) Interface(key string, i interface{}) *LogEvent {
	e.event.Interface(key, i)
	return e
}

// IPAddr TODO
func (e *LogEvent) IPAddr(key string, ip net.IP) *LogEvent {
	e.event.IPAddr(key, ip)
	return e
}

// IPPrefix TODO
func (e *LogEvent) IPPrefix(key string, pfx net.IPNet) *LogEvent {
	e.event.IPPrefix(key, pfx)
	return e
}

// MACAddr TODO
func (e *LogEvent) MACAddr(key string, ha net.HardwareAddr) *LogEvent {
	e.event.MACAddr(key, ha)
	return e
}
