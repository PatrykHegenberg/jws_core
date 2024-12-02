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

func (c *Config) Init(configPath string) error {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %s", err)
	}

	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	return nil
}

func (c *Config) GetLogLevel() string {
	return c.LogLevel
}

// ... weitere Getter-Methoden für andere Felder

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}
	if err := cfg.Init(configPath); err != nil {
		return nil, err
	}
	return cfg, nil
}
