package config

import (
	"github.com/spf13/viper"
	"fmt"
)

type Config struct {
	Mongodb MongoConfig
	Discord DiscordConfig
}

type DiscordConfig struct {
	Token    string
	Username string
	Password string
}

type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

func InitConfig() Config {
	var config Config

	// Search for config.yml in current path
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.AllSettings()

	err = viper.Unmarshal(&config)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Invalid config file: %s \n", err))
	}

	return config
}
