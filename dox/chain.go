package dox

import "context"

// Then ...
func (r Run) Then(o Run) Run {
	return func() error {
		err := r()
		if err != nil {
			return err
		}
		return o()
	}
}

// After ...
func (r Run) After(o Run) Run {
	return func() error {
		err := o()
		if err != nil {
			return err
		}
		return r()
	}
}

// Then ...
func (r RunWithContext) Then(o RunWithContext) RunWithContext {
	return func(ctx context.Context) error {
		err := r(ctx)
		if err != nil {
			return err
		}
		return o(ctx)
	}
}

// After ...
func (r RunWithContext) After(o RunWithContext) RunWithContext {
	return func(ctx context.Context) error {
		err := o(ctx)
		if err != nil {
			return err
		}
		return r(ctx)
	}
}

// Then ...
func (d Do) Then(o Do) Do {
	return func() {
		d()
		o()
	}
}

// After ...
func (d Do) After(o Do) Do {
	return func() {
		o()
		d()
	}
}

// Then ...
func (d DoWithContext) Then(o DoWithContext) DoWithContext {
	return func(ctx context.Context) {
		d(ctx)
		o(ctx)
	}
}

// After ...
func (d DoWithContext) After(o DoWithContext) DoWithContext {
	return func(ctx context.Context) {
		o(ctx)
		d(ctx)
	}
}

//XXX 补充 ThenDo ThenRun ThenDoWithContext ThenRunWithContext ……
