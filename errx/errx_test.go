package errx

import (
	"errors"
	"testing"
)

func TestError_WithData(t *testing.T) {
	err := Code(1).Error("Jack", "a boy")
	e := err.WithMsg("Bob")
	e2 := NewSimpleError(2, "Jack")

	t.Log(e, err)

	t.Log(errors.Is(err, e))
	t.Log(errors.Is(err, e2))
}
