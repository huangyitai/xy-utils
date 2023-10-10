package bindx

import (
	"context"
	"sync/atomic"

	"github.com/huangyitai/xy-utils/dox"
)

// BinderDefAdaptor 将绑定器定义转换为Binder接口的适配器
type BinderDefAdaptor struct {
	def BinderDef
}

// Order ...
func (b *BinderDefAdaptor) Order() int {
	return b.def.Order
}

// Bindable ...
func (b *BinderDefAdaptor) Bindable(slot Slot) bool {
	return b.def.Bindable(slot)
}

// Cacheable ...
func (b *BinderDefAdaptor) Cacheable() bool {
	return b.def.Cacheable()
}

// Bind ...
func (b *BinderDefAdaptor) Bind(ctx context.Context, s ValueSource, slot Slot) error {
	return b.def.Bind(ctx, s, slot)
}

// WatchBinderDefAdaptor 将带监听的绑定器转换为WatchBinder接口的适配器
type WatchBinderDefAdaptor struct {
	def WatchBinderDef
}

// Order ...
func (b *WatchBinderDefAdaptor) Order() int {
	return b.def.Order
}

// Bindable ...
func (b *WatchBinderDefAdaptor) Bindable(slot Slot) bool {
	return b.def.Bindable(slot)
}

// Cacheable ...
func (b *WatchBinderDefAdaptor) Cacheable() bool {
	return b.def.Cacheable()
}

// Bind ...
func (b *WatchBinderDefAdaptor) Bind(ctx context.Context, s ValueSource, slot Slot) error {
	return b.def.Bind(ctx, s, slot)
}

// Watch ...
func (b *WatchBinderDefAdaptor) Watch(ctx context.Context, s ValueSource, slot Slot, callback dox.Run,
	value *atomic.Value) error {
	return b.def.Watch(ctx, s, slot, callback, value)
}
