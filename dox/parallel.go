package dox

import (
	"context"
	"sync"
)

type waitGroupKey struct{}

// While 并发执行，通过while连接的两个函数 a,b 效果上会并发执行，并且最后会等待所有函数执行完成再返回前一个函数 a 的返回值
func (r RunWithContext) While(o RunWithContext) RunWithContext {
	return func(ctx context.Context) error {
		itf := ctx.Value(waitGroupKey{})
		if itf == nil {
			//主协程(invoke发生的协程)
			wg := &sync.WaitGroup{}
			ctx = context.WithValue(ctx, waitGroupKey{}, wg)
			wg.Add(1)

			//拉起子协程
			go func() {
				defer wg.Done()
				_ = o(ctx)
			}()

			//主协程本身继续执行前序逻辑
			err := r(ctx)

			//等待子协程完成
			wg.Wait()
			return err
		} else {
			//主协程（invoke发生的协程）
			wg := itf.(*sync.WaitGroup)
			wg.Add(1)

			//拉起子协程
			go func() {
				defer wg.Done()
				_ = o(ctx)
			}()

			//主协程本身继续执行前序逻辑
			return r(ctx)
		}
	}
}
