package dox

import "time"

var globalWait = NewWaitForCloseSignal()

// DefaultRunBeforeCloseTimeout 默认结束前执行钩子逻辑的超时时间
var DefaultRunBeforeCloseTimeout = 30 * time.Second

// BeforeClose ...
func BeforeClose() *Wait {
	return globalWait
}

// SyncDoBeforeClose ...
func SyncDoBeforeClose(d Do) {
	globalWait.SyncDo(d)
}

// SyncRunBeforeClose ...
func SyncRunBeforeClose(r Run) {
	globalWait.SyncRun(r)
}

// SyncDoWithContextBeforeClose ...
func SyncDoWithContextBeforeClose(d DoWithContext) {
	globalWait.SyncDoWithContext(d)
}

// SyncRunWithContextBeforeClose ...
func SyncRunWithContextBeforeClose(r RunWithContext) {
	globalWait.SyncRunWithContext(r)
}

// AsyncDoBeforeClose ...
func AsyncDoBeforeClose(d Do) {
	globalWait.AsyncDo(d)
}

// AsyncRunBeforeClose ...
func AsyncRunBeforeClose(r Run) {
	globalWait.AsyncRun(r)
}

// AsyncDoWithContextBeforeClose ...
func AsyncDoWithContextBeforeClose(d DoWithContext) {
	globalWait.AsyncDoWithContext(d)
}

// AsyncRunWithContextBeforeClose ...
func AsyncRunWithContextBeforeClose(r RunWithContext) {
	globalWait.AsyncRunWithContext(r)
}

// WaitForCloseSignal 等待结束信号，执行钩子逻辑
func WaitForCloseSignal() error {
	return globalWait.WaitToRunWithTimeout(DefaultRunBeforeCloseTimeout)
}

// InterruptWaitForCloseSignal 中断等待结束信号（钩子逻辑将由其他等待结束信号的协程执行）
func InterruptWaitForCloseSignal() {
	globalWait.Interrupt()
}

// InterruptOrWaitForCloseSignal 中断或等待结束信号，并执行钩子逻辑
func InterruptOrWaitForCloseSignal(interrupt bool) error {
	return globalWait.InterruptOrWaitToRunWithTimeout(DefaultRunBeforeCloseTimeout, interrupt)
}
