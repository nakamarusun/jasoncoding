package config

import (
	"log"
	"github.com/spf13/viper"
)

var config *viper.Viper

// All the necessary variables for this program
var necessaryVariables = []string{
	"GCAPTCHA_SECRET",
	"CONTACT",
}

func checkRequiredVars(config *viper.Viper) {
	// Store all the unloaded variables string
	unloaded := make([]string, 0, len(necessaryVariables))

	for _, vars := range necessaryVariables {
		if config.Get(vars) == nil {
			unloaded = append(unloaded, vars)
		}
	}

	if len(unloaded) != 0 {
		log.Fatalf("Cannot start program! These variables are not set (Lack %d): %v", len(unloaded), unloaded)
	}
}

func Init() {
	config = viper.New()
	config.SetDefault("ENVIRONMENT", "development")
	config.SetDefault("PORT", 8080)
	config.AutomaticEnv()
	checkRequiredVars(config)
}

func GetConfig() *viper.Viper {
	return config
}
