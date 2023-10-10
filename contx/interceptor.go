package contx

import "context"

// ContextRunner ...
type ContextRunner func(ctx context.Context) error

// BeforeIntercept ...
type BeforeIntercept func(ctx context.Context) (context.Context, error)

// Interceptor ...
type Interceptor func(ctx context.Context, r ContextRunner) error

// AfterIntercept ...
type AfterIntercept func(ctx context.Context, err error) error

// NewInterceptBefore ...
func NewInterceptBefore(i BeforeIntercept) Interceptor {
	return func(ctx context.Context, r ContextRunner) error {
		ctx, err := i(ctx)
		if err != nil {
			return err
		}
		return r(ctx)
	}
}

// NewInterceptAfter ...
func NewInterceptAfter(i AfterIntercept) Interceptor {
	return func(ctx context.Context, r ContextRunner) error {
		err := r(ctx)
		return i(ctx, err)
	}
}
