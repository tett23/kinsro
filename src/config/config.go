package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/filesystem"
)

// Config config
type Config struct {
	Environment  string
	VIndexPath   string
	StoragePaths []string
}

var config Config

func envPath() (string, error) {
	goEnv := os.Getenv("GO_ENV")
	dotenvPath, err := dotenvPath(goEnv)
	if err != nil {
		return "", errors.Wrap(err, "envPath")
	}

	return dotenvPath, nil
}

func dotenvPath(environment string) (string, error) {
	if dotenv := os.Getenv("DOTENV_PATH"); dotenv != "" {
		_, err := os.Stat(dotenv)
		if err != nil {
			return "", err
		}

		return dotenv, nil
	}

	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fs := filesystem.GetRawFs()
	envFile := envFilename(environment)

	for currentPath != "/" {
		path := filepath.Join(currentPath, envFile)
		_, err = fs.Stat(path)
		if err == nil {
			return path, nil
		}

		currentPath = filepath.Dir(currentPath)
	}

	return "", errors.Errorf("dotenv file not found. env=%v", environment)
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
