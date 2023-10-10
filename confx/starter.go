package confx

import (
	"context"

	"github.com/huangyitai/xy-utils/contx"
	"github.com/huangyitai/xy-utils/dox"
)

// Starter ...
type Starter struct {
	run dox.RunWithContext
}

// NewStarter ...
func NewStarter() *Starter {
	return &Starter{run: dox.NoopRunWithContext}
}

// Start ...
func (s *Starter) Start(ctx context.Context, r contx.ContextRunner) error {
	err := s.run(ctx)
	if err != nil {
		return err
	}
	err = r(ctx)
	return err
}

// SetDefault ...
func (s *Starter) SetDefault(read ReadFunc) *Starter {
	s.run = s.run.Then(func(ctx context.Context) error {
		Default = read
		return nil
	})
	return s
}

// EnableBind ...
func (s *Starter) EnableBind() *Starter {
	s.run = s.run.Then(func(ctx context.Context) error {
		EnableBind()
		return nil
	})
	return s
}
