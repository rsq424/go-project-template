package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	System System
}

var Cfg Config

func GetConfig() *Config {
	return &Cfg
}

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(".")
	vp.AutomaticEnv()
	err := vp.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func Setup() error {
	setting, err := NewSetting()
	if err != nil {
		return err
	}

	_ = setting.vp.Unmarshal(&Cfg.System)
	return nil
}
