package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
	Port          string `mapstructure:"PORT"`
}

func LoadConfig() {
	err := LoadConfigFromOs()
	if err == nil {
		return
	}

	log.Println(err.Error())

	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()

	if err != nil {
		panic(err.Error())
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Printf("Using config %+v", config)
	log.Printf("Using port %s", config.Port)
	log.Printf("Using weather api key %s", config.WeatherApiKey)
}

func LoadConfigFromOs() error {
	config.WeatherApiKey = os.Getenv("WEATHER_API_KEY")
	if config.WeatherApiKey == "" {
		return errors.New("WEATHER_API_KEY must be set")
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		return errors.New("PORT must be set")
	}

	return nil
}

func GetConfig() Config {
	return config
}
