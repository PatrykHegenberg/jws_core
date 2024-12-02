package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel  string `mapstructure:"log_level"`
	LogFormat string `mapstructure:"log_format"`
	LogFile   string `mapstructure:"log_file"`
}

var cfg Config

func InitConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	return nil
}

func GetConfig() Config {
	return cfg
}
