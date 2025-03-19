package config

type Database struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}