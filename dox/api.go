package dox

import "context"

// NoopDo ...
func NoopDo() {}

// NoopDoWithContext ...
func NoopDoWithContext(ctx context.Context) {}

// NoopRun ...
func NoopRun() error {
	return nil
}

// NoopRunWithContext ...
func NoopRunWithContext(ctx context.Context) error {
	return nil
}
