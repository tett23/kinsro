package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// Config config
type Config struct {
	Environment  string
	VIndexPath   string
	StoragePaths []string
}

var config Config

func envPath() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "aaa", err
	}

	goEnv := os.Getenv("GO_ENV")
	// dotenvPath, err := filepath.Abs(filepath.Join(pwd, "../../", envFilename((goEnv))))
	dotenvPath, err := filepath.Abs(filepath.Join(pwd, envFilename((goEnv))))
	if err != nil {
		return "bbb", err
	}

	return dotenvPath, nil
}

func envFilename(environment string) string {
	switch environment {
	case "development":
		return ".development.env"
	case "test":
		return ".test.env"
	case "production":
		return ".production.env"
	case "":
		return ".env"
	default:
		panic(fmt.Sprintf("Unexpected environment name. %v", environment))
	}
}

func init() {
	dotenvPath, err := envPath()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(dotenvPath)
	if err != nil {
		panic(err)
	}

	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		fmt.Println("Incorrect GO_ENV")
	}

	vindexPath := os.Getenv("VINDEX_PATH")
	if vindexPath == "" {
		fmt.Println("Incorrect VINDEX_PATH")
	}

	storagePaths := strings.Split(os.Getenv("STORAGE_PATHS"), ",")

	config = Config{
		Environment:  goEnv,
		VIndexPath:   vindexPath,
		StoragePaths: storagePaths,
	}
}

// GetConfig GetConfig
func GetConfig() *Config {
	return &config
}
