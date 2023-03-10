package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	// viper读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	getString := viper.GetString("obj")
	host := viper.GetString("mysql.host")
	fmt.Println(getString)
	fmt.Println("host:", host)
}

func TestReadConfigWithYaml(t *testing.T) {
	config := &Config{}
	// 读取config文件
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}
	fmt.Printf("config: %+v", config)
}
