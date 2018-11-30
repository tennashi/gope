package config

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Base represent the base config.
type Base struct {
	*viper.Viper
}

// NewBaseConfig construct a new config.
func NewBaseConfig(fileName string) *Base {
	config := &Base{viper.New()}
	setDefaultConfig(config)
	setConfigByFlags(config)

	config.WatchConfig()
	config.SetConfigFile(fileName)
	config.SetConfigType("toml")

	if err := config.ReadInConfig(); err != nil {
		log.Println(err)
		log.Println("use default config")
	}
	return config
}

func setDefaultConfig(config *Base) {
	config.SetDefault("port", "1323")
	config.SetDefault("project_file", "./projects.toml")
}

func setConfigByFlags(config *Base) {
	config.BindPFlag("port", pflag.Lookup("port"))
	config.BindPFlag("project_file", pflag.Lookup("project"))
}
