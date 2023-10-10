package bindx

import (
	"context"
	"reflect"
	"sync/atomic"

	"github.com/huangyitai/xy-utils/dox"
)

// SubTypeBindable ...
type SubTypeBindable struct {
	reflect.Type
}

// Bindable ...
func (c *SubTypeBindable) Bindable(slot Slot) bool {
	return slot.Type().AssignableTo(c.Type)
}

// TypeBindable ...
type TypeBindable struct {
	reflect.Type
}

// Bindable ...
func (c *TypeBindable) Bindable(slot Slot) bool {
	return slot.Type() == c.Type
}

// SuperTypeBindable ...
type SuperTypeBindable struct {
	reflect.Type
}

// Bindable ...
func (c *SuperTypeBindable) Bindable(slot Slot) bool {
	return c.Type.AssignableTo(slot.Type())
}

type zeroBinder struct{}

// Bindable ...
func (z zeroBinder) Bindable(slot Slot) bool {
	return false
}

// Cacheable ...
func (z zeroBinder) Cacheable() bool {
	return true
}

// Bind ...
func (z zeroBinder) Bind(ctx context.Context, s ValueSource, slot Slot) error {
	zero := reflect.Zero(slot.Value().Type())
	slot.Value().Set(zero)
	return nil
}

type noopBinder struct{}

// Watch ...
func (n noopBinder) Watch(ctx context.Context, s ValueSource, slot Slot, callback dox.Run, value *atomic.Value) error {
	t := slot.Type().Out(0)
	v, err := s.GetValue(ctx, slot.Name(), t, string(slot.Tag()))
	if err != nil {
		return err
	}
	value.Store(v.Interface())
	return callback()
}

// Bindable ...
func (n noopBinder) Bindable(slot Slot) bool {
	return false
}

// Cacheable ...
func (n noopBinder) Cacheable() bool {
	return true
}

// Bind ...
func (n noopBinder) Bind(ctx context.Context, s ValueSource, slot Slot) error {
	return nil
}
