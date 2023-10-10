package dox

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func sleepAndPrint(duration time.Duration) {
	time.Sleep(duration)
	fmt.Println("Done!")
}

func TestDo_WithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	err := Do(func() {
		sleepAndPrint(time.Second * 1)
	}).WithContext(ctx)()
	t.Log(err)
}

func TestDo_WithContextAndTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	err := Do(func() {
		sleepAndPrint(time.Second * 4)
	}).WithContextAndTimeout(ctx, time.Second*3)()
	t.Log(err)
}

func TestDo_WithTimeout(t *testing.T) {
	err := Do(func() {
		sleepAndPrint(time.Second * 3)
	}).WithTimeout(time.Second * 4)()
	t.Log(err)
}

func TestDo_PanicToErr(t *testing.T) {
	err := Do(func() {
		panic(1)
	}).PanicToErr()()
	t.Log(err)

	Do(func() {
		sleepAndPrint(time.Second)
		panic(1)
	}).Go()

}
