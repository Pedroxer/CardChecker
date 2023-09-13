package util

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
}

func LoadConfig(path string) (conf Config, err error) {
	viper.SetConfigName("Config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("Cannot load config: %s", err)
	}
	err = viper.Unmarshal(&conf)
	return
}
