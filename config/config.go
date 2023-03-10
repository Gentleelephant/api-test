package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var ApiTestConfig *Config

type Config struct {
	Obj   string `json:"obj" yaml:"obj"`
	Mysql Mysql  `json:"mysql" yaml:"mysql"`
}

type Mysql struct {
	Host string `json:"host" yaml:"host"`

	Port string `json:"port" yaml:"port"`

	Username string `json:"username" yaml:"username"`

	Password string `json:"password" yaml:"password"`
}

// LoadConfig 加载配置文件
func LoadConfig() {
	config := new(Config)
	file, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		panic("read file error")
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic("unmarshal file error")
	}
	ApiTestConfig = config
}
