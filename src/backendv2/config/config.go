package config

import (
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)

var Cfg *viper.Viper

// All the necessary variables for this program
var necessaryVariables = []string{
	"GCAPTCHA_SECRET",
	"CONTACT",
	"JWT_SECRET",
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
	Cfg = viper.New()
	cwd, err := os.Getwd()
	if err == nil {
		Cfg.SetDefault("FONT_PATH", path.Join(cwd, "/fonts"))
	}

	Cfg.SetDefault("ENVIRONMENT", "development")
	Cfg.SetDefault("PORT", 8080)
	Cfg.AutomaticEnv()
	checkRequiredVars(Cfg)
}
