package metricx

type noopAPI struct{}

// Counter ...
func (n noopAPI) Counter(metric string, value float64, tags map[string]string) error {
	return nil
}

// Gauge ...
func (n noopAPI) Gauge(metric string, value float64, tags map[string]string) error {
	return nil
}

// AvgGauge ...
func (n noopAPI) AvgGauge(metric string, value float64, tags map[string]string) error {
	return nil
}

// MaxCounter ...
func (n noopAPI) MaxCounter(metric string, value float64, tags map[string]string) error {
	return nil
}

// MinCounter ...
func (n noopAPI) MinCounter(metric string, value float64, tags map[string]string) error {
	return nil
}

// Incr ...
func (n noopAPI) Incr(metric string, tags map[string]string) error {
	return nil
}

// Event ...
func (n noopAPI) Event(strMark, msg string, tags map[string]string) error {
	return nil
}

// Histogram ...
func (n noopAPI) Histogram(metric string, value float64, tags map[string]string, buckets []float64) error {
	return nil
}

// EventWithReceiver ...
func (n noopAPI) EventWithReceiver(strMark, msg string, receivers []string, tags map[string]string) error {
	return nil
}
