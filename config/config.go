package config

import (
	"fmt"

	"strings"

	"github.com/jeff3710/ndot/pkg/log"
	"github.com/spf13/viper"
)

type Config struct {
	App        App               `mapstructure:"app" yaml:"app"`
	Database   Database          `mapstructure:"database" yaml:"database"`
	Vendors    map[string]string `mapstructure:"vendors" yaml:"vendors"`
	SystemOIDs map[string]string `mapstructure:"system_oids" yaml:"system_oids"`
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
		log.Errorf("配置加载失败详情",
			log.String("config_file", viper.ConfigFileUsed()),
		)
		return nil, fmt.Errorf("配置加载失败: %w", err)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {

		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &config, nil
}
