package dox

import (
	"context"
	"io"
	"time"

	"github.com/huangyitai/xy-utils/deferx"
)

// Try ...
type Try func(attempt int) (retry bool, err error)

// TryWithContext ...
type TryWithContext func(ctx context.Context, attempt int) (retry bool, err error)

// Do ...
type Do func()

// DoWithContext ...
type DoWithContext func(context.Context)

// Run ...
type Run func() error

// RunWithContext ...
type RunWithContext func(context.Context) error

// WithContext ...
func (d Do) WithContext(ctx context.Context) Run {
	return func() error {
		end := make(chan bool, 1)
		go func() {
			d()
			end <- true
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-end:
			return nil
		}
	}
}

// ContextAware ...
func (d Do) ContextAware() RunWithContext {
	return func(ctx context.Context) error {
		end := make(chan bool, 1)
		go func() {
			d()
			end <- true
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-end:
			return nil
		}
	}
}

// AutoClose ...
func (d Do) AutoClose(closer io.Closer) Run {
	return func() (err error) {
		defer deferx.Close(closer, &err)
		d()
		return
	}
}

// WithContextAndTimeout ...
func (d Do) WithContextAndTimeout(ctx context.Context, timeout time.Duration) Run {
	return func() error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return d.WithContext(ctx)()
	}
}

// WithContextAndTimeoutNS ...
func (d Do) WithContextAndTimeoutNS(ctx context.Context, timeout int64) Run {
	return d.WithContextAndTimeout(ctx, time.Duration(timeout))
}

// WithContextAndDeadline ...
func (d Do) WithContextAndDeadline(ctx context.Context, deadline time.Time) Run {
	return func() error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return d.WithContext(ctx)()
	}
}

// WithTimeout ...
func (d Do) WithTimeout(timeout time.Duration) Run {
	return d.WithContextAndTimeout(context.TODO(), timeout)
}

// WithTimeoutNS ...
func (d Do) WithTimeoutNS(timeout int64) Run {
	return d.WithTimeout(time.Duration(timeout))
}

// WithDeadline ...
func (d Do) WithDeadline(deadline time.Time) Run {
	return d.WithContextAndDeadline(context.TODO(), deadline)
}

// PanicToErr ...
func (d Do) PanicToErr() Run {
	return func() (err error) {
		defer deferx.PanicToError(&err)
		d()
		return
	}
}

// AutoClose ...
func (d DoWithContext) AutoClose(closer io.Closer) RunWithContext {
	return func(ctx context.Context) (err error) {
		defer deferx.Close(closer, &err)
		d(ctx)
		return
	}
}

// SetTimeout ...
func (d DoWithContext) SetTimeout(timeout time.Duration) DoWithContext {
	return func(ctx context.Context) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		d(ctx)
	}
}

// WithTimeout ...
func (d DoWithContext) WithTimeout(timeout time.Duration) RunWithContext {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return Do(func() {
			d(ctx)
		}).WithContext(ctx)()
	}
}

// WithTimeoutNS ...
func (d DoWithContext) WithTimeoutNS(timeout int64) RunWithContext {
	return d.WithTimeout(time.Duration(timeout))
}

// WithDeadline ...
func (d DoWithContext) WithDeadline(deadline time.Time) RunWithContext {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return Do(func() {
			d(ctx)
		}).WithContext(ctx)()
	}
}

// PanicToErr ...
func (d DoWithContext) PanicToErr() RunWithContext {
	return func(ctx context.Context) (err error) {
		defer deferx.PanicToError(&err)
		d(ctx)
		return
	}
}

// WithRetry ...
func (r Run) WithRetry(maxRetries int) Run {
	return func() (err error) {
		retries := 0
		for maxRetries < 0 || retries <= maxRetries {
			err = r()
			if err == nil {
				break
			}
			retries++
		}
		return err
	}
}

// AutoClose ...
func (r Run) AutoClose(closer io.Closer) Run {
	return func() (err error) {
		defer deferx.Close(closer, &err)
		err = r()
		return
	}
}

// WithContext ...
func (r Run) WithContext(ctx context.Context) Run {
	return func() error {
		end := make(chan error, 1)
		go func() {
			end <- r()
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-end:
			return err
		}
	}
}

// ContextAware ...
func (r Run) ContextAware() RunWithContext {
	return func(ctx context.Context) error {
		end := make(chan error, 1)
		go func() {
			end <- r()
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-end:
			return err
		}
	}
}

// WithContextAndTimeout ...
func (r Run) WithContextAndTimeout(ctx context.Context, timeout time.Duration) Run {
	return func() error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return r.WithContext(ctx)()
	}
}

// WithContextAndTimeoutNS ...
func (r Run) WithContextAndTimeoutNS(ctx context.Context, timeout int64) Run {
	return r.WithContextAndTimeout(ctx, time.Duration(timeout))
}

// WithContextAndDeadline ...
func (r Run) WithContextAndDeadline(ctx context.Context, deadline time.Time) Run {
	return func() error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return r.WithContext(ctx)()
	}
}

// WithTimeout ...
func (r Run) WithTimeout(timeout time.Duration) Run {
	return r.WithContextAndTimeout(context.TODO(), timeout)
}

// WithTimeoutNS ...
func (r Run) WithTimeoutNS(timeout int64) Run {
	return r.WithTimeout(time.Duration(timeout))
}

// WithDeadline ...
func (r Run) WithDeadline(deadline time.Time) Run {
	return r.WithContextAndDeadline(context.TODO(), deadline)
}

// PanicToErr ...
func (r Run) PanicToErr() Run {
	return func() (err error) {
		defer deferx.PanicToError(&err)
		return r()
	}
}

// IgnoreErr ...
func (r Run) IgnoreErr() Do {
	return func() {
		_ = r()
	}
}

// WithRetry ...
func (r RunWithContext) WithRetry(maxRetries int) RunWithContext {
	return func(ctx context.Context) (err error) {
		retries := 0
		for maxRetries < 0 || retries <= maxRetries {
			err = r(ctx)
			if err == nil {
				break
			}
			retries++
		}
		return err
	}
}

// AutoClose ...
func (r RunWithContext) AutoClose(closer io.Closer) RunWithContext {
	return func(ctx context.Context) (err error) {
		defer deferx.Close(closer, &err)
		err = r(ctx)
		return
	}
}

// SetTimeout ...
func (r RunWithContext) SetTimeout(timeout time.Duration) RunWithContext {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return r(ctx)
	}
}

// WithTimeout ...
func (r RunWithContext) WithTimeout(timeout time.Duration) RunWithContext {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return Run(func() error {
			return r(ctx)
		}).WithContext(ctx)()
	}
}

// WithTimeoutNS ...
func (r RunWithContext) WithTimeoutNS(timeout int64) RunWithContext {
	return r.WithTimeout(time.Duration(timeout))
}

// WithDeadline ...
func (r RunWithContext) WithDeadline(deadline time.Time) RunWithContext {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return Run(func() error {
			return r(ctx)
		}).WithContext(ctx)()
	}
}

// PanicToErr ...
func (r RunWithContext) PanicToErr() RunWithContext {
	return func(ctx context.Context) (err error) {
		defer deferx.PanicToError(&err)
		return r(ctx)
	}
}

// IgnoreErr ...
func (r RunWithContext) IgnoreErr() DoWithContext {
	return func(ctx context.Context) {
		_ = r(ctx)
	}
}
