package server

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	User         string
	Password     string
	Host         string
	Port         int
	DBName       string
	SslMode      string
	MaxOpenConns int
	MaxIdleConns int
}

// LogConfig 定义日志配置结构体
type LogConfig struct {
	DisableCaller     bool     `mapstructure:"disable-caller"`
	DisableStacktrace bool     `mapstructure:"disable-stacktrace"`
	Level             string   `mapstructure:"level"`
	Format            string   `mapstructure:"format"`
	OutputPaths       []string `mapstructure:"output-paths"`
}

type Config struct {
	App        AppConfig         `mapstructure:"app"`
	Database   DatabaseConfig    `mapstructure:"database"`
	Log        LogConfig         `mapstructure:"log"`
	Vendors    map[string]string `mapstructure:"vendors"`
	SystemOIDs map[string]string `mapstructure:"system_oids"`
}

func LoadConfig() (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.Debug()
	viper.AddConfigPath(".")

	// 支持环境变量覆盖配置
	viper.AutomaticEnv()
	viper.SetEnvPrefix("NDOT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	return &config, nil
}
