package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string
	// ...
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		ServerPort: viper.GetString("server.port"),
		// ...
	}

	return config, nil
}
