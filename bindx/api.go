package bindx

import (
	"context"
	"reflect"
	"sync/atomic"

	"github.com/huangyitai/xy-utils/dox"
)

// Default ...
var Default = NewContainer()

// Binder ...
type Binder interface {
	Bindable(slot Slot) bool
	Cacheable() bool
	Bind(ctx context.Context, s ValueSource, slot Slot) error
}

// WatchBinder ...
type WatchBinder interface {
	Binder
	Watch(ctx context.Context, s ValueSource, slot Slot, callback dox.Run, value *atomic.Value) error
}

// OrderedBinder ...
type OrderedBinder interface {
	Order() int
	Binder
}

// OrderedWatchBinder ...
type OrderedWatchBinder interface {
	Order() int
	WatchBinder
}

// ValueSource ...
type ValueSource interface {
	BindOne(ctx context.Context, name string, i interface{}, tags ...string) error
	GetValue(ctx context.Context, name string, t reflect.Type, tags ...string) (reflect.Value, error)
	WatchOne(ctx context.Context, name string, i interface{}, callback dox.Run, tags ...string) error
}

// Slot ...
type Slot interface {
	Tag() reflect.StructTag
	Name() string
	Value() reflect.Value
	Type() reflect.Type
	WatchType() reflect.Type
	Get() interface{}
	Set(i interface{})
	Pointer() interface{}
	SetPointer(i interface{})
	SetValue(v reflect.Value)
	SetIndirectValue(v reflect.Value)
	Order() int
}

// NewContainer ...
func NewContainer() *Container {
	return &Container{
		bindingTable: map[string]map[reflect.Type]*binding{},
		typeConv:     map[string]map[reflect.Type]reflect.Type{},
		binderSet:    map[Binder]bool{},
		binders:      []OrderedBinder{},
		slots:        []Slot{},
		bindings:     []*binding{},
	}
}

// AddOne ...
func AddOne(name string, i interface{}, tags ...string) error {
	return Default.AddOne(name, i, tags...)
}

// AddCombo ...
func AddCombo(a ...interface{}) error {
	return Default.AddCombo(a...)
}

// AddBinder ...
func AddBinder(binder Binder) {
	Default.AddBinder(binder)
}

// Bind ...
func Bind(ctx context.Context) (err error) {
	return Default.Bind(ctx)
}

// BindOne ...
func BindOne(ctx context.Context, name string, i interface{}, tags ...string) error {
	return Default.BindOne(ctx, name, i, tags...)
}

// GetValue ...
func GetValue(ctx context.Context, name string, t reflect.Type, tags ...string) (reflect.Value, error) {
	return Default.GetValue(ctx, name, t, tags...)
}
