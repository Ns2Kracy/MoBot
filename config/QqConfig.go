package config

type QqConfig struct {
	Qq       int64  `yaml:"qq"`
	Password string `yaml:"password"`
	Login    bool   `yaml:"login"`
}
