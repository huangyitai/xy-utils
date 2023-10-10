package bindx

import (
	"context"
	"sync/atomic"

	"github.com/huangyitai/xy-utils/dox"
)

// BinderDef 绑定器定义，可以通过设置各个函数字段实现绑定器
type BinderDef struct {
	Order     int
	Bindable  func(slot Slot) bool
	Cacheable func() bool
	Bind      func(ctx context.Context, s ValueSource, slot Slot) error
}

// Binder ...
func (d BinderDef) Binder() OrderedBinder {
	return &BinderDefAdaptor{def: d}
}

// WatchBinderDef 带监听的绑定器定义，可以通过设置各个函数字段实现带监听的绑定器
type WatchBinderDef struct {
	BinderDef
	Watch func(ctx context.Context, s ValueSource, slot Slot, callback dox.Run, value *atomic.Value) error
}

// Binder ...
func (d WatchBinderDef) Binder() OrderedWatchBinder {
	return &WatchBinderDefAdaptor{def: d}
}
