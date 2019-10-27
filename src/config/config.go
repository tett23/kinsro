package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/filesystem"
)

// Config config
type Config struct {
	Environment      string
	VIndexPath       string
	StoragePaths     []string
	FfmpegScriptPath string
	FfmpegPresetPath string
}

var config Config

// Initialize Initialize
func Initialize() {
	if os.Getenv("GO_ENV") == "test" {
		filesystem.CopyDotfile()
	}

	envConfig := NewPartialConfigFromEnvironmentVariable()
	if envConfig.Fulfillded() {
		conf, err := envConfig.ToConfig()
		if err != nil {
			panic(err)
		}

		config = *conf

		return
	}

	fs := filesystem.GetFs()
	dotenvPath, ok, err := findDotenvPath(fs)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic(errors.Errorf("dotenv file not found"))
	}

	fileConfig, err := NewPartialConfigFromFile(fs, dotenvPath)
	if err != nil {
		panic(err)
	}

	merged := fileConfig.Merge(envConfig)
	conf, err := merged.ToConfig()
	if err != nil {
		panic(err)
	}

	config = *conf
}

// GetConfig GetConfig
func GetConfig() *Config {
	return &config
}
