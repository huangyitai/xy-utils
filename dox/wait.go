package dox

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Wait ...
type Wait struct {
	syncRun  []RunWithContext
	asyncRun []RunWithContext
	ready    chan interface{}
	once     sync.Once
}

// NewWait ...
func NewWait(waitFunc func()) *Wait {
	w := &Wait{
		ready: make(chan interface{}, 1),
	}

	go func() {
		defer func() {
			w.ready <- nil
		}()
		waitFunc()
	}()
	return w
}

// NewWaitForCloseSignal ...
func NewWaitForCloseSignal() *Wait {
	return NewWait(func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)
		<-ch
	})
}

// SyncDo ...
func (w *Wait) SyncDo(d Do) {
	w.syncRun = append(w.syncRun, d.ContextAware())
}

// AsyncDo ...
func (w *Wait) AsyncDo(d Do) {
	w.asyncRun = append(w.asyncRun, d.ContextAware())
}

// SyncRun ...
func (w *Wait) SyncRun(r Run) {
	w.syncRun = append(w.syncRun, r.ContextAware())
}

// AsyncRun ...
func (w *Wait) AsyncRun(r Run) {
	w.asyncRun = append(w.asyncRun, r.ContextAware())
}

// SyncDoWithContext ...
func (w *Wait) SyncDoWithContext(d DoWithContext) {
	w.syncRun = append(w.syncRun, d.PanicToErr())
}

// AsyncDoWithContext ...
func (w *Wait) AsyncDoWithContext(d DoWithContext) {
	w.asyncRun = append(w.asyncRun, d.PanicToErr())
}

// SyncRunWithContext ...
func (w *Wait) SyncRunWithContext(r RunWithContext) {
	w.syncRun = append(w.syncRun, r)
}

// AsyncRunWithContext ...
func (w *Wait) AsyncRunWithContext(r RunWithContext) {
	w.asyncRun = append(w.asyncRun, r)
}

func (w *Wait) run(ctx context.Context) error {
	var err error
	for _, run := range w.syncRun {
		err = run(ctx)
		if err != nil {
			return err
		}
	}

	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(w.asyncRun))
	for _, run := range w.asyncRun {
		r := run
		go func() {
			defer func() {
				wg.Done()
			}()
			e := r(ctx)
			if e != nil {
				m.Lock()
				err = e
				m.Unlock()
			}
		}()
	}
	wg.Wait()
	return err
}

// WaitToRunWithContext ...
func (w *Wait) WaitToRunWithContext(ctx context.Context) (err error) {
	w.once.Do(func() {
		<-w.ready
		err = w.run(ctx)
	})
	return
}

// WaitToRunWithTimeout ...
func (w *Wait) WaitToRunWithTimeout(timeout time.Duration) (err error) {
	w.once.Do(func() {
		<-w.ready
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err = w.run(ctx)
	})
	return
}

// WaitToRun ...
func (w *Wait) WaitToRun() (err error) {
	w.once.Do(func() {
		<-w.ready
		err = w.run(context.Background())
	})
	return
}

// Interrupt ...
func (w *Wait) Interrupt() {
	select {
	case w.ready <- nil:
	default:
	}
}

// InterruptOrWaitToRun ...
func (w *Wait) InterruptOrWaitToRun(interrupt bool) error {
	if interrupt {
		w.Interrupt()
	}
	return w.WaitToRun()
}

// InterruptOrWaitToRunWithTimeout ...
func (w *Wait) InterruptOrWaitToRunWithTimeout(timeout time.Duration, interrupt bool) error {
	if interrupt {
		w.Interrupt()
	}
	return w.WaitToRunWithTimeout(timeout)
}

// InterruptOrWaitToRunWithContext ...
func (w *Wait) InterruptOrWaitToRunWithContext(ctx context.Context, interrupt bool) error {
	if interrupt {
		w.Interrupt()
	}
	return w.WaitToRunWithContext(ctx)
}
