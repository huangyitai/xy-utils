package tracex

// Event ...
type Event struct {
	api API
}

// Tag ...
func (e *Event) Tag(key string, value string) *Event {
	e.api.Tag(key, value)
	return e
}

// Log ...
func (e *Event) Log(kvp ...string) *Event {
	e.api.Log(kvp...)
	return e
}
