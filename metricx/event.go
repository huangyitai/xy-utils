package metricx

// Event ...
type Event struct {
	api  API
	tags map[string]string
	ops  []func(*Event) error
}

// Tag ...
func (e *Event) Tag(name, value string) *Event {
	e.tags[name] = value
	return e
}

// Tags ...
func (e *Event) Tags(tags map[string]string) *Event {
	for k, v := range tags {
		e.tags[k] = v
	}
	return e
}

// Counter ...
func (e *Event) Counter(metric string, value float64) *Event {
	e.ops = append(e.ops, func(e *Event) error {
		return e.api.Counter(metric, value, e.tags)
	})
	return e
}

// Gauge ...
func (e *Event) Gauge(metric string, value float64) *Event {
	e.ops = append(e.ops, func(e *Event) error {
		return e.api.Gauge(metric, value, e.tags)
	})
	return e
}

// AvgGauge ...
func (e *Event) AvgGauge(metric string, value float64) *Event {
	e.ops = append(e.ops, func(e *Event) error {
		return e.api.AvgGauge(metric, value, e.tags)
	})
	return e
}

// MaxCounter ...
func (e *Event) MaxCounter(metric string, value float64) *Event {
	e.ops = append(e.ops, func(e *Event) error {
		return e.api.MaxCounter(metric, value, e.tags)
	})
	return e
}

// MinCounter ...
func (e *Event) MinCounter(metric string, value float64) *Event {
	e.ops = append(e.ops, func(e *Event) error {
		return e.api.MinCounter(metric, value, e.tags)
	})
	return e
}

// Incr ...
func (e *Event) Incr(metric string) *Event {
	e.ops = append(e.ops, func(e *Event) error {
		return e.api.Incr(metric, e.tags)
	})
	return e
}

// Histogram ...
func (e *Event) Histogram(metric string, value float64, buckets []float64) *Event {
	e.ops = append(e.ops, func(event *Event) error {
		return e.api.Histogram(metric, value, e.tags, buckets)
	})
	return e
}

// Event ...
func (e *Event) Event(strMark, msg string) *Event {
	e.ops = append(e.ops, func(e *Event) error {
		return e.api.Event(strMark, msg, e.tags)
	})
	return e
}

// EventWithReceiver ...
func (e *Event) EventWithReceiver(strMark, msg string, receivers []string) *Event {
	e.ops = append(e.ops, func(e *Event) error {
		return e.api.EventWithReceiver(strMark, msg, receivers, e.tags)
	})
	return e
}

// Send ...
func (e *Event) Send() error {
	for _, op := range e.ops {
		err := op(e)
		if err != nil {
			return err
		}
	}
	return nil
}
