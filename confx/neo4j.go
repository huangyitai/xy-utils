package confx

type baseNeo4jConfig struct {
	//Host 主机名/ip
	Host string

	//Port 端口号
	Port int

	//User 用户名
	User string

	//Password 密码
	Password string

	//DBName 数据库名称
	DBName string

	//MaxIdleConns 连接池参数：最大空闲连接数
	MaxIdleConns int

	//MaxOpenConns 连接池参数：最大连接数
	MaxOpenConns int

	//ConnMaxLifetime 连接池参数：连接被重用的最大时长，超时的连接会在再次被重用之前关闭，单位分钟
	ConnMaxLifetime int

	//ConnectTimeout 尝试与mysql服务器连接时，服务器返回错误前等待客户端数据包的最大时长，例如：30s、0.5m、1m30s
	ConnectTimeout string
}

// Neo4jGoConfig TODO
type Neo4jGoConfig struct {
	ExtraProps
	BindDefault
	baseNeo4jConfig `yaml:",inline" mapstructure:",squash"`
}
