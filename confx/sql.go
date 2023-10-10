package confx

// BaseSQLConfig 基本sql数据库配置
type BaseSQLConfig struct {
	ExtraProps
	BindDefault

	// Host 主机名/ip
	Host string

	// Port 端口号
	Port int

	// DBName 数据库名称
	DBName string
}
