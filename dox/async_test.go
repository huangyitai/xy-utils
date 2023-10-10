package dox

import (
	"context"
	"fmt"
	"testing"
)

func TestRun_Go(t *testing.T) {
	end := Do(func() {
		panic(nil)
	}).Go()
	t.Log("end", <-end)
}

func TestRunWithContext_Daemon(t *testing.T) {
	cnt := 0

	end := RunWithContext(func(ctx context.Context) error {
		t.Log(cnt)
		cnt++
		if cnt < 5 {
			panic(nil)
		}
		return fmt.Errorf("exit")
	}).Daemon(context.Background())

	t.Log("end", <-end)
}

func TestRun_Daemon(t *testing.T) {
	cnt := 0
	end := Run(func() error {
		t.Log(cnt)
		cnt++
		if cnt < 15 {
			panic(nil)
		}
		return fmt.Errorf("exit")
	}).Daemon()
	t.Log("end", <-end)
}
