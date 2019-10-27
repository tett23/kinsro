package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

func findDotenvPath(fs afero.Fs) (string, bool, error) {
	envPath := os.Getenv("DOTENV_PATH")
	dotenv, ok, err := checkEnvironmentVariable(fs, envPath)
	if err != nil {
		return "", false, err
	}
	if ok {
		return dotenv, true, nil
	}

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "production"
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", false, err
	}

	dotenv, ok, err = checkCurrentDirectory(fs, env, pwd)
	if err != nil {
		return "", false, err
	}
	if ok {
		return dotenv, true, nil
	}

	home := os.Getenv("HOME")
	if home == "" {
		home = "/"
	}

	dotenv, ok, err = checkConfigDirectory(fs, env, home)
	if err != nil {
		return "", false, err
	}
	if ok {
		return dotenv, true, nil
	}

	return "", false, errors.Errorf("dotenv file not found.")
}

func checkEnvironmentVariable(fs afero.Fs, path string) (string, bool, error) {
	if path == "" {
		return "", false, nil
	}

	if !isExist(fs, path) {
		return "", false, nil
	}

	return path, true, nil
}

func checkCurrentDirectory(fs afero.Fs, env, pwd string) (string, bool, error) {
	return checkRecursive(fs, env, pwd)
}

// ConfigDir ConfigDir
const ConfigDir = "$HOME/.config/kinsro"

func checkConfigDirectory(fs afero.Fs, env, home string) (string, bool, error) {
	dir, err := expandHomeDirectory(ConfigDir, home)
	if err != nil {
		return "", false, err
	}

	return checkDir(fs, env, dir)
}

func expandHomeDirectory(path, home string) (string, error) {
	ret := strings.ReplaceAll(ConfigDir, "$HOME", home)
	ret = strings.ReplaceAll(ret, "~/", home+"/")

	return filepath.Abs(ret)
}

func checkRecursive(fs afero.Fs, env, path string) (string, bool, error) {
	if path == "/" {
		return checkDir(fs, env, path)
	}

	ret, ok, err := checkDir(fs, env, path)
	if err != nil {
		return "", false, err
	}
	if ok {
		return ret, true, nil
	}

	return checkRecursive(fs, env, filepath.Dir(path))
}

func checkDir(fs afero.Fs, env, path string) (string, bool, error) {
	expected := []string{
		filepath.Join(path, ".env."+env),
		filepath.Join(path, ".env"),
	}

	for i := range expected {
		if isExist(fs, expected[i]) {
			return expected[i], true, nil
		}
	}

	return "", false, nil
}

func isExist(fs afero.Fs, path string) bool {
	_, err := fs.Stat(path)

	return err == nil
}
