package config

import (
	"os"
)

var config Config

type Config struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
	Port          string `mapstructure:"PORT"`
}

func LoadConfig() {
	config.WeatherApiKey = os.Getenv("WEATHER_API_KEY")
	if config.WeatherApiKey == "" {
		panic("WEATHER_API_KEY must be set")
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		panic("PORT must be set")
	}
}

func GetConfig() Config {
	return config
}
