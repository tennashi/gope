package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config is wrapper
type Config struct {
	*viper.Viper
}

const configName = "config"

// NewConfig is constructor of Config
func NewConfig(filePath string) *Config {
	config := &Config{viper.New()}
	setDefaultConfig(config)
	config.WatchConfig()
	config.SetConfigName(configName)
	config.SetConfigType("toml")
	config.AddConfigPath(filePath)
	if err := config.ReadInConfig(); err != nil {
		log.Println(err)
		log.Println("use default config")
	}
	return config
}

func setDefaultConfig(config *Config) {
	config.SetDefault("base_path", ".")
	config.SetDefault("port", "1323")
}
