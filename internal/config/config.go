package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	AppEnv              = "APP_ENV"
	InfuraURI           = "INFURA_URI"
	InfuraProjectID     = "INFURA_PROJECT_ID"
	InfuraProjectSecret = "INFURA_PROJECT_SECRET"
	ServerPort          = "SERVER_PORT"
	ContractAddress     = "CONTRACT_ADDRESS"
)

const (
	Development = "development"
	Production  = "production"
)

type Config struct {
	AppEnv              string
	ServerPort          string
	InfuraURI           string
	InfuraProjectID     string
	InfuraProjectSecret string
	ContractAddress     string
}

func New() *Config {
	c := new(Config)
	c.Load()

	return c
}

func (c *Config) Load() {
	var appEnv string
	viper.AutomaticEnv()
	viper.SetDefault(AppEnv, Development)

	appEnv = viper.Get(AppEnv).(string)
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigName(appEnv)
	viper.SetConfigType("env")
	viper.AddConfigPath(fmt.Sprintf("%s/internal/config", root))
	viper.AddConfigPath(fmt.Sprintf("..%s/internal/config", root))
	viper.AddConfigPath(fmt.Sprintf("%s/../config", root))
	viper.AddConfigPath(".")

	if err = viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	c.AppEnv = appEnv
	c.ServerPort = viper.Get(ServerPort).(string)
	c.InfuraURI = viper.Get(InfuraURI).(string)
	c.InfuraProjectID = viper.Get(InfuraProjectID).(string)
	c.InfuraProjectSecret = viper.Get(InfuraProjectSecret).(string)
	c.ContractAddress = viper.Get(ContractAddress).(string)

	fmt.Println(fmt.Sprintf("%s environment loaded successfully !", c.AppEnv))
}
