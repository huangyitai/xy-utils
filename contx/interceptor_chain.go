package contx

import (
	"context"
	"sort"
)

// InterceptorChain 拦截器链，将按照slice顺序嵌套执行
type InterceptorChain []Interceptor

// Intercept ...
func (ic InterceptorChain) Intercept(ctx context.Context, r ContextRunner) error {
	if len(ic) == 0 {
		return r(ctx)
	}

	return ic[0](ctx, func(ctx context.Context) error {
		return ic[1:].Intercept(ctx, r)
	})
}

// Intercept ...
func Intercept(ctx context.Context, r ContextRunner, interceptors ...Interceptor) error {
	return InterceptorChain(interceptors).Intercept(ctx, r)
}

// OrderedInterceptor 有序拦截器，包含Order信息，可以根据Order顺序进行排序得到有序的拦截器链
type OrderedInterceptor struct {
	Interceptor
	Order int
}

// GetOrderedInterceptorChain 按照Order构造有序拦截器链
func GetOrderedInterceptorChain(ascending bool, its ...*OrderedInterceptor) InterceptorChain {
	//先进行slice拷贝，避免修改传入的slice
	itsCopy := make([]*OrderedInterceptor, len(its))
	copy(itsCopy, its)

	sort.SliceStable(itsCopy, func(i, j int) bool {
		if ascending {
			return itsCopy[i].Order < itsCopy[j].Order
		} else {
			return itsCopy[i].Order > itsCopy[j].Order
		}
	})

	is := make([]Interceptor, len(itsCopy))
	for i, it := range itsCopy {
		is[i] = it.Interceptor
	}
	return is
}

// GetOrderedInterceptorChainAsc 按Order升序构造拦截器链
func GetOrderedInterceptorChainAsc(its ...*OrderedInterceptor) InterceptorChain {
	return GetOrderedInterceptorChain(true, its...)
}

// GetOrderedInterceptorChainDesc 按Order降序构造拦截器链
func GetOrderedInterceptorChainDesc(its ...*OrderedInterceptor) InterceptorChain {
	return GetOrderedInterceptorChain(false, its...)
}
