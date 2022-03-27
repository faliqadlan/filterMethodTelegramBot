package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Token string
}

var synchronizer = &sync.Mutex{}

var appConfig *AppConfig

func initConfig() *AppConfig {
	config := AppConfig{}
	config.Token = os.Getenv("TOKEN")

	return &config
}

func defaultConfig() *AppConfig {
	config := AppConfig{}

	if err := godotenv.Load("local.env"); err != nil {
		log.Warn(err)
		return &config
	}

	config.Token = os.Getenv("TOKEN")

	return &config
}

func GetConfig() *AppConfig {
	synchronizer.Lock()
	defer synchronizer.Unlock()

	appConfig = initConfig()

	if appConfig.Token == "" {
		appConfig = defaultConfig()
	}

	return appConfig
}
