package config

import (
	"log"

	"github.com/spf13/viper"
	"jasoncoding.com/backendv2/utils"
)

var config *viper.Viper

func Init() {
	env := utils.GetEnvDefault("ENVIRONMENT", "development")

	config = viper.New()
	config.SetConfigType("env")
	config.SetConfigName(env)
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error loading configuration file '%s' (%v)", env, err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
