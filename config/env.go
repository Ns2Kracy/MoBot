package config

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                   // 环境值
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                //地址
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                // 端口
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`       // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"` // 使用redis
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
