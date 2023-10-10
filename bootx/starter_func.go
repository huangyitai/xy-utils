package bootx

import (
	"context"

	"github.com/huangyitai/xy-utils/contx"
)

// StarterFunc 单函数启动器，相当于只提供Start函数的实现
type StarterFunc func(ctx context.Context, r contx.ContextRunner) error

// Start 直接调用函数本身
func (s StarterFunc) Start(ctx context.Context, r contx.ContextRunner) error {
	return s(ctx, r)
}
