package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port      string
	DbHost    string
	DbPort    string
	DbUser    string
	DbPass    string
	DbName    string
	JwtSecret string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		Port:      viper.GetString("PORT"),
		DbHost:    viper.GetString("DB_HOST"),
		DbPort:    viper.GetString("DB_PORT"),
		DbUser:    viper.GetString("DB_USER"),
		DbPass:    viper.GetString("DB_PASSWORD"),
		DbName:    viper.GetString("DB_NAME"),
		JwtSecret: viper.GetString("JWT_SECRET"),
	}

	return config, nil
}