package config

import (
	"currencyParser/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	MainDatabase config.MainDatabase
	Cache        config.Cache
	Log          config.Log
	Telegram     config.Telegram
	Binance      config.Binance
	Exmo         config.Exmo
	isInited     bool
}

var configInstance Config

func init() {
	if !configInstance.isInited {
		configInstance = Config{}
	}
	err := configInstance.init()
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return &configInstance
}

func (config *Config) init() error {
	var err error

	err = envconfig.Init(&config.MainDatabase)
	if err != nil {
		return err
	}

	err = envconfig.Init(&config.Cache)
	if err != nil {
		return err
	}

	err = envconfig.Init(&config.Log)
	if err != nil {
		return err
	}

	err = envconfig.Init(&config.Telegram)
	if err != nil {
		return err
	}

	err = envconfig.Init(&config.Binance)
	if err != nil {
		return err
	}

	err = envconfig.Init(&config.Exmo)
	if err != nil {
		return err
	}

	config.isInited = true

	return err
}