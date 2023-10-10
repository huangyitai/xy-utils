package metricx

import (
	"context"
	"testing"
)

func TestExample(t *testing.T) {
	ctx := context.TODO()

	//请求来了，先在ctx里初始化好维度
	ctx = New().Tag("主调信息", "xxxxxx").Tag("被调信息", "xxxxxx").WithContext(ctx)

	//具体上报某个指标
	_ = Ctx(ctx).Report().
		Incr("xxxxx").Send()

}
