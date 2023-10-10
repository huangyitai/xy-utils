package confx

import (
	"context"
	"reflect"
	"sync/atomic"

	"github.com/huangyitai/xy-utils/bindx"
	"github.com/huangyitai/xy-utils/dox"
	"github.com/huangyitai/xy-utils/tagx"
	"github.com/huangyitai/xy-utils/xxx"
	"github.com/rs/zerolog/log"
)

// EnableBind ...
func EnableBind() {
	bindx.AddBinder(binder)
}

// BindDefault ...
type BindDefault struct{}

// WatchFunc ...
func (b BindDefault) WatchFunc() WatchFunc {
	return DefaultWatch
}

// ReadFunc ...
func (b BindDefault) ReadFunc() ReadFunc {
	return Default
}

// BindConfig ...
type BindConfig struct{}

// WatchFunc ...
func (b BindConfig) WatchFunc() WatchFunc {
	return nil
}

// ReadFunc ...
func (b BindConfig) ReadFunc() ReadFunc {
	return nil
}

// Binding ...
type Binding interface {
	ReadFunc() ReadFunc
	WatchFunc() WatchFunc
}

// NeedInitDefault 需要在读取配置前初始化默认值（设置的默认值可以被用户指定的空值覆盖）
type NeedInitDefault interface {
	InitDefault()
}

// NeedSetDefault 需要在读取配置后填充默认值（用户未指定或指定为空值会被设置为默认值）
type NeedSetDefault interface {
	SetDefault()
}

var bindingType = reflect.TypeOf((*Binding)(nil)).Elem()

var binder = Binder{}

// Binder ...
type Binder struct{}

// Bindable ...
func (b Binder) Bindable(slot bindx.Slot) bool {
	if !bindableKind(slot) {
		return false
	}

	if slot.Type().AssignableTo(bindingType) {
		return true
	}

	info := TagInfo{}
	if tagx.UnmarshalTag(TagKey, &info, slot.Tag()) != nil {
		return false
	}
	return info.Binding
}

func isUnsupportedKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Chan, reflect.Func, reflect.Invalid, reflect.Uintptr, reflect.UnsafePointer, reflect.Interface:
		return true
	default:
		return false
	}
}

func bindableKind(slot bindx.Slot) bool {
	//暂时只支持 结构体 或 结构体的指针

	if isUnsupportedKind(slot.Type().Kind()) {
		return false
	}

	if slot.Type().Kind() == reflect.Ptr && isUnsupportedKind(slot.Type().Elem().Kind()) {
		return false
	}

	return true
}

// Cacheable ...
func (b Binder) Cacheable() bool {
	return true
}

// Bind ...
func (b Binder) Bind(ctx context.Context, s bindx.ValueSource, slot bindx.Slot) error {
	sign := xxx.NewSignStr().WithPath("confx", "binder").WithProp(bindx.TagKey, slot.Name())
	log.Trace().
		Str("sName", slot.Name()).
		Str("sTag", string(slot.Tag())).
		Msgf("%s bind start", sign)

	if slot.Value().Kind() == reflect.Ptr && slot.Value().IsZero() {
		slot.SetValue(reflect.New(slot.Type().Elem()))
	}

	err := ReadField(slot.Name(), slot.Value(), slot.Tag())
	if err != nil {
		log.Err(err).Msgf("%s ReadField fail", sign)
		return err
	}
	log.Trace().Str("sName", slot.Name()).Msgf("%s bind end", sign)
	return nil
}

// Watch ...
func (b Binder) Watch(ctx context.Context, s bindx.ValueSource, slot bindx.Slot, callback dox.Run,
	value *atomic.Value) error {
	sign := xxx.NewSignStr().WithPath("confx", "binder").WithProp(bindx.TagKey, slot.Name())
	log.Trace().
		Str("sName", slot.Name()).
		Str("sTag", string(slot.Tag())).
		Str("sType", slot.Type().String()).
		Msgf("%s watch start", sign)

	return WatchField(slot.Name(), slot.WatchType(), func(v reflect.Value) error {
		value.Store(v.Interface())
		return callback()
	}, slot.Tag())
}
