package contx

import (
	"context"
	"time"
)

type detachedContext struct{ parent context.Context }

// Detach 返回脱离控制的Context, 屏蔽ctx中包含的超时、取消等控制
func Detach(ctx context.Context) context.Context { return detachedContext{ctx} }

// Deadline implements context.Deadline
func (v detachedContext) Deadline() (time.Time, bool) { return time.Time{}, false }

// Done implements context.Done
func (v detachedContext) Done() <-chan struct{} { return nil }

// Err implements context.Err
func (v detachedContext) Err() error { return nil }

// Value implements context.Value
func (v detachedContext) Value(key interface{}) interface{} { return v.parent.Value(key) }
