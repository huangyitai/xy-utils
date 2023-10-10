package dox

import (
	"context"
	"github.com/pkg/errors"
)

// Go ...
func (d Do) Go() chan error {
	return Run(func() error {
		d()
		return nil
	}).Go()
}

// Go ...
func (r Run) Go() chan error {
	ch := make(chan error, 1)
	go func() {
		ok := false
		defer func() {
			if !ok {
				i := recover()
				if err, ok := i.(error); ok {
					ch <- err
				} else {
					ch <- errors.Errorf("%+v", i)
				}
			}
		}()

		ch <- r()
		ok = true
	}()
	return ch
}

// Go ...
func (d DoWithContext) Go(ctx context.Context) chan error {
	return RunWithContext(func(ctx context.Context) error {
		d(ctx)
		return nil
	}).Go(ctx)
}

// Go ...
func (r RunWithContext) Go(ctx context.Context) chan error {
	ch := make(chan error, 1)
	go func() {
		ok := false
		defer func() {
			if !ok {
				i := recover()
				if err, ok := i.(error); ok {
					ch <- err
				} else {
					ch <- errors.Errorf("%+v", i)
				}
			}
		}()

		ch <- r(ctx)
		ok = true
	}()
	return ch
}

// Daemon ...
func (r RunWithContext) Daemon(ctx context.Context) chan error {
	res := make(chan error, 1)
	ch := make(chan error, 1)
	go func() {
		restart := true
		var e error

		for restart {
			restart = false

			go func() {
				ok := false
				defer func() {
					if !ok {
						restart = true
						i := recover()
						if err, ok := i.(error); ok {
							ch <- err
						} else {
							ch <- errors.Errorf("%+v", i)
						}
					}
				}()
				ch <- r(ctx)
				ok = true
			}()

			e = <-ch
		}

		res <- e
	}()
	return res
}

// Daemon ...
func (d DoWithContext) Daemon(ctx context.Context) chan error {
	return RunWithContext(func(ctx context.Context) error {
		d(ctx)
		return nil
	}).Daemon(ctx)
}

// Daemon ...
func (d Do) Daemon() chan error {
	return Run(func() error {
		d()
		return nil
	}).Daemon()
}

// Daemon ...
func (r Run) Daemon() chan error {
	res := make(chan error, 1)
	ch := make(chan error, 1)
	go func() {
		restart := true
		var e error

		for restart {
			restart = false

			go func() {
				ok := false
				defer func() {
					if !ok {
						restart = true
						i := recover()
						if err, ok := i.(error); ok {
							ch <- err
						} else {
							ch <- errors.Errorf("%+v", i)
						}
					}
				}()
				ch <- r()
				ok = true
			}()

			e = <-ch
		}

		res <- e
	}()
	return res
}
