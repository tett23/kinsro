package config

import (
	"bytes"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// PartialConfig PartialConfig
type PartialConfig struct {
	Environment  *string
	VIndexPath   *string
	StoragePaths *[]string
}

// Fulfillded Fulfillded
func (partialConfig PartialConfig) Fulfillded() bool {
	return partialConfig.Environment != nil && partialConfig.VIndexPath != nil && partialConfig.StoragePaths != nil
}

// Merge Merge
func (partialConfig PartialConfig) Merge(updater *PartialConfig) *PartialConfig {
	ret := PartialConfig{
		Environment:  partialConfig.Environment,
		VIndexPath:   partialConfig.VIndexPath,
		StoragePaths: partialConfig.StoragePaths,
	}

	if updater.Environment != nil {
		ret.Environment = updater.Environment
	}
	if updater.VIndexPath != nil {
		ret.VIndexPath = updater.VIndexPath
	}
	if updater.Environment != nil {
		ret.StoragePaths = updater.StoragePaths
	}

	return &ret
}

// ToConfig ToConfig
func (partialConfig PartialConfig) ToConfig() (*Config, error) {
	ret := Config{}

	if partialConfig.Environment == nil {
		return nil, errors.Errorf("PartialConfig.ToConfig: cast error. field=Environment")
	}
	ret.Environment = *partialConfig.Environment

	if partialConfig.VIndexPath == nil {
		return nil, errors.Errorf("PartialConfig.ToConfig: cast error. field=VIndexPath")
	}
	ret.VIndexPath = *partialConfig.VIndexPath

	if partialConfig.StoragePaths == nil {
		return nil, errors.Errorf("PartialConfig.ToConfig: cast error. field=StoragePaths")
	}
	ret.StoragePaths = *partialConfig.StoragePaths

	return &ret, nil
}

// NewPartialConfigFromEnvironmentVariable NewPartialConfigFromEnvironmentVariable
func NewPartialConfigFromEnvironmentVariable() *PartialConfig {
	partial := PartialConfig{}

	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		goEnv = "development"
	}
	partial.Environment = &goEnv

	var vindexPath string = os.Getenv("VINDEX_PATH")
	if vindexPath == "" {
		partial.VIndexPath = nil
	} else {
		partial.VIndexPath = &vindexPath
	}

	storagePaths := strings.Split(os.Getenv("STORAGE_PATHS"), ",")
	if len(storagePaths) == 0 {
		partial.StoragePaths = nil
	} else {
		partial.StoragePaths = &storagePaths
	}

	return &partial
}

// NewPartialConfigFromFile NewPartialConfigFromFile
func NewPartialConfigFromFile(fs afero.Fs, path string) (*PartialConfig, error) {
	if !isExist(fs, path) {
		return nil, errors.Errorf("dotfile does not exists. path=%v", path)
	}

	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(data)
	envMap, err := godotenv.Parse(buf)
	if err != nil {
		return nil, err
	}

	for k, v := range envMap {
		os.Setenv(k, v)
	}

	return NewPartialConfigFromEnvironmentVariable(), nil
}
