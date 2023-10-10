package skywalking

// Config skywalking配置信息
type Config struct {
	Service          string            `yaml:"service" mapstructure:"service"`                         // 实例名 默认为服务名
	Address          string            `yaml:"address" mapstructure:"address"`                         //skywalking服务器名称 ip：port
	CheckInterval    string            `yaml:"check_interval" mapstructure:"check_interval"`           //conn 健康检查时间间隔
	MaxSendQueueSize int               `yaml:"max_send_queue_size" mapstructure:"max_send_queue_size"` // 可发送消息队列
	Auth             string            `yaml:"auth" mapstructure:"auth"`                               // skywalking 鉴权信息
	InstanceProps    map[string]string `yaml:"props" mapstructure:"props"`                             // 元属性
	ComponentId      int32             `yaml:"component_id" mapstructure:"component_id"`               // 组件id
	SamplingRate     float64           `yaml:"sampling_rate" mapstructure:"sampling_rate"`             // 采样率
}

// NewConfig 创建新的配置信息
func NewConfig() *Config {
	return &Config{
		SamplingRate: 1.0,
	}
}
