package config

import (
	"github.com/spf13/viper"
	"log"
)

type configuration struct {
	Server   server
	Logger   logger
	InMemory inMemory
	DataBase Database
}

type server struct {
	Port     int
	DataType string
}

type inMemory struct {
	Path string
}

type Database struct {
	Path      string
	TableName string
}

var (
	instance   *configuration
	configPath = "./configs/"
)

func GetInstance() *configuration {
	if instance == nil {
		instance = newConfig()
	}
	return instance
}

func newConfig() *configuration {
	var conf configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error read config! %v", err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Error read config! %v", err)
	}
	conf.Logger.newLogger()
	return &conf
}
