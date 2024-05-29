package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddr string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func LoadConfig() error {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return &Config{
		ServerAddr: viper.GetString("SERVER_ADDR"),
		DbHost:     viper.GetString("DB_HOST"),
		DbPort:     viper.GetString("DB_PORT"),
		DbName:     viper.GetString("DB_NAME"),
		DbUser:     viper.GetString("DB_USER"),
		DbPassword: viper.GetString("DB_PASSWORD"),
	}
}
