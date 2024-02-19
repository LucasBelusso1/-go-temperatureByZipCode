package config

import "github.com/spf13/viper"

var Config *viper.Viper

func LoadConfig(path string) error {
	Config = viper.New()
	Config.SetConfigName("app_config")
	Config.SetConfigType("env")
	Config.AddConfigPath(path)
	Config.SetConfigFile(".env")
	Config.AutomaticEnv()

	err := Config.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return nil
}
