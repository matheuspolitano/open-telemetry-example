package config

import (
	"github.com/spf13/viper"
)

type Server struct {
	Port string
}

type Config struct {
	Server Server
}

func LoadConfig(pathConfig string) (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName(pathConfig)
	vp.AddConfigPath(".")
	vp.AutomaticEnv()
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return vp, nil
}
