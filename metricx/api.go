package metricx

import (
	"context"
)

// Default ...
var Default API = noopAPI{}

// APIs ...
var APIs = map[string]API{}

// API ...
type API interface {
	Counter(metric string, value float64, tags map[string]string) error
	Gauge(metric string, value float64, tags map[string]string) error
	AvgGauge(metric string, value float64, tags map[string]string) error
	MaxCounter(metric string, value float64, tags map[string]string) error
	MinCounter(metric string, value float64, tags map[string]string) error
	Incr(metric string, tags map[string]string) error
	Histogram(metric string, value float64, tags map[string]string, buckets []float64) error
	Event(strMark, msg string, tags map[string]string) error
	EventWithReceiver(strMark, msg string, receivers []string, tags map[string]string) error
}

// GetAPI ...
func GetAPI(name string) API {
	api := APIs[name]
	if api == nil {
		return noopAPI{}
	}
	return api
}

// New ...
func New() *Context {
	return &Context{
		api:  Default,
		tags: map[string]string{},
	}
}

// NewByName ...
func NewByName(name string) *Context {
	return &Context{
		api:  GetAPI(name),
		tags: map[string]string{},
	}
}

// Ctx ...
func Ctx(ctx context.Context) *Context {
	i := ctx.Value(contextKey{})
	if i != nil {
		if c, ok := i.(*Context); ok {
			return c
		}
	}

	return &Context{api: noopAPI{}, tags: map[string]string{}}
}

// Report ...
func Report() *Event {
	return &Event{
		api:  Default,
		tags: map[string]string{},
	}
}

// ReportByName ...
func ReportByName(name string) *Event {
	return &Event{
		api:  GetAPI(name),
		tags: map[string]string{},
	}
}
