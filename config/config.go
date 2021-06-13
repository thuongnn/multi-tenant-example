package config

import (
	"example/src/models"
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error read in config .env")
	}
}

func DatabaseConfig() *models.Database {
	return &models.Database{
		Host:         viper.GetString("DB_HOST"),
		Port:         viper.GetInt("DB_PORT"),
		Username:     viper.GetString("DB_USER"),
		Password:     viper.GetString("DB_PASSWORD"),
		Database:     viper.GetString("DB_NAME"),
	}
}