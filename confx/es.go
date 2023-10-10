package confx

// baseConfig elastic基本配置信息
type baseESConfig struct {
	// es地址：比如http://127.0.0.1:9200
	URL string
	// es用户账号名
	Username string
	// es密码
	Password string
}

// ES7Config es7.0版本专用配置
type ES7Config struct {
	BindDefault
	baseESConfig `yaml:",inline" mapstructure:",squash"`
}
