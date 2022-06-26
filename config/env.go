package config

type System struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"` //地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	WsPort   string `mapstructure:"ws_port" json:"ws_port" yaml:"ws_port"`
	DbType   string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`       // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	UseRedis bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"` // 使用redis
}
