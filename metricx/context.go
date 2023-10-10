package metricx

import "context"

type contextKey struct{}

// Context ...
type Context struct {
	api  API
	tags map[string]string
}

// Tag ...
func (c *Context) Tag(name, value string) *Context {
	c.tags[name] = value
	return c
}

// Tags ...
func (c *Context) Tags(tags map[string]string) *Context {
	for k, v := range tags {
		c.tags[k] = v
	}
	return c
}

// WithContext ...
func (c *Context) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey{}, c)
}

// Report ...
func (c *Context) Report() *Event {
	tags := map[string]string{}
	for k, v := range c.tags {
		tags[k] = v
	}
	return &Event{
		api:  c.api,
		tags: tags,
	}
}

// ReportByName ...
func (c *Context) ReportByName(name string) *Event {
	tags := map[string]string{}
	for k, v := range c.tags {
		tags[k] = v
	}
	return &Event{
		api:  GetAPI(name),
		tags: tags,
	}
}

// Clone ...
func (c *Context) Clone() *Context {
	tags := map[string]string{}
	for k, v := range c.tags {
		tags[k] = v
	}
	return &Context{
		api:  c.api,
		tags: tags,
	}
}
