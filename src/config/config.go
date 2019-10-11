package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config config
type Config struct {
	Environment string
	VIndexPath  string
}

var config Config

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		log.Fatal("Incorrect GO_ENV")
	}

	vindexPath := os.Getenv("VINDEX_PATH")
	if vindexPath == "" {
		log.Fatal("Incorrect VINDEX_PATH")
	}

	config = Config{
		Environment: goEnv,
		VIndexPath:  vindexPath,
	}
}

// GetConfig GetConfig
func GetConfig() *Config {
	return &config
}
