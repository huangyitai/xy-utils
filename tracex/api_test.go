package tracex

import (
	"context"
	"testing"
)

func TestCtx(t *testing.T) {
	ctx := context.Background()
	Ctx(ctx).Tag("123", "456")
	Ctx(ctx).Log("123", "456")
}
