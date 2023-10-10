package confx

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"sync"
	"sync/atomic"
)

var hooks = []mapstructure.DecodeHookFunc{
	mapstructure.StringToTimeDurationHookFunc(),
	mapstructure.StringToSliceHookFunc(","),
}

var hookLock sync.Mutex
var hookOption atomic.Value

func defaultDecoderConfigOptions() []viper.DecoderConfigOption {
	i := hookOption.Load()
	if i == nil {
		return []viper.DecoderConfigOption{}
	} else {
		return []viper.DecoderConfigOption{i.(func(c *mapstructure.DecoderConfig))}
	}
}

// AddTypeHook ...
func AddTypeHook(hook mapstructure.DecodeHookFuncType) {
	if hook == nil {
		return
	}

	hookLock.Lock()
	defer hookLock.Unlock()

	hooks = append(hooks, hook)
	hookOption.Store(func(c *mapstructure.DecoderConfig) {
		c.DecodeHook = mapstructure.ComposeDecodeHookFunc(hooks...)
	})
}

// AddKindHook ...
func AddKindHook(hook mapstructure.DecodeHookFuncKind) {
	if hook == nil {
		return
	}

	hookLock.Lock()
	defer hookLock.Unlock()

	hooks = append(hooks, hook)
	hookOption.Store(func(c *mapstructure.DecoderConfig) {
		c.DecodeHook = mapstructure.ComposeDecodeHookFunc(hooks...)
	})
}
