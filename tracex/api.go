package tracex

import "context"

// API ...
type API interface {
	Tag(string, string)
	Log(...string)
	ID() string
}

type noopAPI struct{}

// ID ...
func (n noopAPI) ID() string {
	return "no_trace"
}

// Tag ...
func (n noopAPI) Tag(s string, s2 string) {
}

// Log ...
func (n noopAPI) Log(s ...string) {
}

type contextKey struct{}

// WithCtx ...
func WithCtx(ctx context.Context, api API) context.Context {
	return context.WithValue(ctx, contextKey{}, api)
}

// Ctx ...
func Ctx(ctx context.Context) *Event {
	itf := ctx.Value(contextKey{})
	if itf == nil {
		return &Event{api: noopAPI{}}
	}

	api, ok := itf.(API)
	if !ok {
		return &Event{api: noopAPI{}}
	}

	return &Event{api: api}
}

// Tag ...
func Tag(ctx context.Context, key, value string) {
	Ctx(ctx).Tag(key, value)
}

// Log ...
func Log(ctx context.Context, kvp ...string) {
	Ctx(ctx).Log(kvp...)
}

// ID ...
func ID(ctx context.Context) string {
	return Ctx(ctx).api.ID()
}
