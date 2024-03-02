package config

import (
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Mode         string
}

type Config struct {
	Server Server
	Logger LoggerConfig
}

type LoggerConfig struct { // Logger config
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
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

func ParseConfig(viperConfig *viper.Viper) (*Config, error) {
	var configParsed Config
	err := viperConfig.Unmarshal(&configParsed)
	if err != nil {
		return nil, err
	}
	return &configParsed, nil

}
