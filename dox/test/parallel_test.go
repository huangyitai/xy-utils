package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/huangyitai/xy-utils/dox"
)

func hello(ctx context.Context) error {
	time.Sleep(time.Second * 1)
	fmt.Println("hello 1!")
	return nil
}

func helloThenPanic(ctx context.Context) error {
	time.Sleep(time.Second * 5)
	fmt.Println("hello 2!")
	panic(1)
}

func TestWhile(t *testing.T) {
	b := dox.RunWithContext(helloThenPanic).PanicToErr()

	err := dox.RunWithContext(hello).
		Then(b).After(b).
		While(b).
		While(b).
		While(b).
		While(b)(context.Background())

	fmt.Println(err)
}
