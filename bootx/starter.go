package bootx

import (
	"context"
	"time"

	"github.com/huangyitai/xy-utils/contx"
)

// Starter ...
type Starter interface {
	Start(ctx context.Context, r contx.ContextRunner) error
}

// CombinedStarter ...
type CombinedStarter struct {
	chain contx.InterceptorChain
}

// Start ...
func (c *CombinedStarter) Start(ctx context.Context, r contx.ContextRunner) error {
	return c.chain.Intercept(ctx, r)
}

// CombineStarters ...
func CombineStarters(starters ...Starter) *CombinedStarter {
	res := CombinedStarter{
		chain: []contx.Interceptor{},
	}

	for _, starter := range starters {
		if starter == nil {
			continue
		}

		res.chain = append(res.chain, starter.Start)
	}
	return &res
}

// TimeoutStarter ...
type TimeoutStarter struct {
	timeout time.Duration
}

// Start ...
func (t *TimeoutStarter) Start(ctx context.Context, r contx.ContextRunner) error {
	ctx, cancel := context.WithTimeout(ctx, t.timeout)
	defer cancel()
	return r(ctx)
}

// WithTimeout ...
func WithTimeout(timeout time.Duration, starters ...Starter) Starter {
	res := CombinedStarter{
		chain: []contx.Interceptor{func(ctx context.Context, r contx.ContextRunner) error {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			return r(ctx)
		}},
	}

	for _, starter := range starters {
		if starter == nil {
			continue
		}

		res.chain = append(res.chain, starter.Start)
	}
	return &res
}
