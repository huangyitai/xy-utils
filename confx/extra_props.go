package confx

import (
	"errors"
	"github.com/spf13/cast"
)

// PropNotFound 未找到对应Property
var PropNotFound = errors.New("prop not found")

// ExtraProps 用于提供额外的自定义参数
type ExtraProps map[string]interface{}

// GetString ...
func (p ExtraProps) GetString(key string) (string, error) {
	if itf, ok := p[key]; ok {
		return cast.ToStringE(itf)
	}
	return "", PropNotFound
}

// GetBool ...
func (p ExtraProps) GetBool(key string) (bool, error) {
	if itf, ok := p[key]; ok {
		return cast.ToBoolE(itf)
	}
	return false, PropNotFound
}

// GetInt ...
func (p ExtraProps) GetInt(key string) (int, error) {
	if itf, ok := p[key]; ok {
		return cast.ToIntE(itf)
	}
	return 0, PropNotFound
}

// GetStringSlice ...
func (p ExtraProps) GetStringSlice(key string) ([]string, error) {
	if itf, ok := p[key]; ok {
		return cast.ToStringSliceE(itf)
	}
	return nil, PropNotFound
}
