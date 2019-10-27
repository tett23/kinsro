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
	Environment      *string
	VIndexPath       *string
	StoragePaths     *[]string
	FfmpegScriptPath *string
	FfmpegPresetPath *string
}

// Fulfillded Fulfillded
func (partialConfig PartialConfig) Fulfillded() bool {
	return partialConfig.Environment != nil &&
		partialConfig.VIndexPath != nil &&
		partialConfig.StoragePaths != nil &&
		partialConfig.FfmpegScriptPath != nil &&
		partialConfig.FfmpegPresetPath != nil
}

// Merge Merge
func (partialConfig PartialConfig) Merge(updater *PartialConfig) *PartialConfig {
	ret := PartialConfig{
		Environment:      partialConfig.Environment,
		VIndexPath:       partialConfig.VIndexPath,
		StoragePaths:     partialConfig.StoragePaths,
		FfmpegScriptPath: partialConfig.FfmpegScriptPath,
		FfmpegPresetPath: partialConfig.FfmpegPresetPath,
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
	if updater.FfmpegScriptPath != nil {
		ret.FfmpegScriptPath = updater.FfmpegScriptPath
	}
	if updater.FfmpegPresetPath != nil {
		ret.FfmpegPresetPath = updater.FfmpegPresetPath
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

	if partialConfig.FfmpegScriptPath == nil {
		return nil, errors.Errorf("PartialConfig.ToConfig: cast error. field=FfmpegScriptPath")
	}
	ret.FfmpegScriptPath = *partialConfig.FfmpegScriptPath

	if partialConfig.FfmpegPresetPath == nil {
		return nil, errors.Errorf("PartialConfig.ToConfig: cast error. field=FfmpegPresetPath")
	}
	ret.FfmpegPresetPath = *partialConfig.FfmpegPresetPath

	return &ret, nil
}

// NewPartialConfigFromEnvironmentVariable NewPartialConfigFromEnvironmentVariable
func NewPartialConfigFromEnvironmentVariable() *PartialConfig {
	partial := PartialConfig{}

	if value := os.Getenv("GO_ENV"); value == "" {
		value = "development"
	} else {
		partial.Environment = &value
	}

	if value := os.Getenv("VINDEX_PATH"); value == "" {
		partial.VIndexPath = nil
	} else {
		partial.VIndexPath = &value
	}

	if value := strings.Split(os.Getenv("STORAGE_PATHS"), ","); len(value) == 0 {
		partial.StoragePaths = nil
	} else {
		partial.StoragePaths = &value
	}

	if value := os.Getenv("FFSHELL_PATH"); value == "" {
		partial.FfmpegScriptPath = nil
	} else {
		partial.FfmpegScriptPath = &value
	}

	if value := os.Getenv("FFPRESET_PATH"); value == "" {
		partial.FfmpegPresetPath = nil
	} else {
		partial.FfmpegPresetPath = &value
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
