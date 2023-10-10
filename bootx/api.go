package bootx

import (
	"context"
)

var noopStart = func(ctx context.Context) error { return nil }

// Run ...
func Run(starters ...Starter) error {
	return RunWithContext(context.Background(), starters...)
}

// JustRun ...
func JustRun(starters ...Starter) {
	_ = Run(starters...)
}

// RunWithContext ...
func RunWithContext(ctx context.Context, starters ...Starter) error {
	starter := CombineStarters(starters...)
	return starter.Start(ctx, noopStart)
}

// JustRunWithContext ...
func JustRunWithContext(ctx context.Context, starters ...Starter) {
	_ = RunWithContext(ctx, starters...)
}
