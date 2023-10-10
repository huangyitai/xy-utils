package contx

import "context"

// CallWithContext 在context的超时、取消控制下，调用闭包，并返回闭包结果
// 如果context超时或被取消时运行仍未完成，则返回错误，但闭包会在另一协程继续执行
func CallWithContext(ctx context.Context, call func() error) error {
	res := make(chan error, 1)
	go func() {
		res <- call()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-res:
		return err
	}
}

// RunWithContext 在context的超时、取消控制下，运行闭包
// 如果context超时或被取消时运行仍未完成，则返回错误，但闭包会在另一协程继续执行
func RunWithContext(ctx context.Context, run func()) error {
	end := make(chan bool, 1)
	go func() {
		run()
		end <- true
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-end:
		return nil
	}
}
