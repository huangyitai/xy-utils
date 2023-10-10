package validatex

import validator "github.com/go-playground/validator/v10"

// Default 工具库维护的默认全局validate对象
var Default = validator.New()

// Get 获取全局validate对象
func Get() *validator.Validate {
	return Default
}
