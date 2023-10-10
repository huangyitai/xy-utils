package dox

import (
	"fmt"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	w := NewWait(func() {
		<-time.After(time.Second * 2)
	})
	w.SyncDo(func() {
		fmt.Println("hello!")
	})
	w.AsyncDo(func() {
		fmt.Println("123")
	})
	w.AsyncDo(func() {
		fmt.Println("333")
	})
	_ = w.WaitToRun()
}
