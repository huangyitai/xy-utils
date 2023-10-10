package bindx

import (
	"context"

	"github.com/huangyitai/xy-utils/contx"
	"github.com/huangyitai/xy-utils/dox"
)

// Starter ...
type Starter struct {
	run dox.RunWithContext
}

// NewStarter TODO
func NewStarter() *Starter {
	return &Starter{run: dox.NoopRunWithContext}
}

// Start ...
func (s *Starter) Start(ctx context.Context, r contx.ContextRunner) error {
	err := s.run.Then(Bind)(ctx)
	if err != nil {
		return err
	}
	err = r(ctx)
	return err
}

// AddBinder ...
func (s *Starter) AddBinder(binder Binder) *Starter {
	s.run = s.run.Then(func(ctx context.Context) error {
		AddBinder(binder)
		return nil
	})
	return s
}

// BindOne ...
func (s *Starter) BindOne(name string, i interface{}, tags ...string) *Starter {
	s.run = s.run.Then(func(ctx context.Context) error {
		return AddOne(name, i, tags...)
	})
	return s
}

// BindCombo ...
func (s *Starter) BindCombo(a ...interface{}) *Starter {
	s.run = s.run.Then(func(ctx context.Context) error {
		return AddCombo(a...)
	})
	return s
}
