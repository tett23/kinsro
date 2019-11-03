package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tett23/kinsro/src/filesystem"
)

func Test__Config__checkEnvironmentVariable(t *testing.T) {
	fs := filesystem.GetFs()
	envPath := "/test/.env"
	defer func() {
		filesystem.ResetTestFs()
	}()

	t.Run("should return nil", func(t *testing.T) {
		_, ok, err := checkEnvironmentVariable(fs, envPath)
		if ok {
			t.Log("fail")
			t.Fail()
		}
		if err != nil {
			t.Log("fail")
			t.Error(err)
		}
	})

	t.Run("should retun string", func(t *testing.T) {
		if _, err := fs.Create(envPath); err != nil {
			t.Log("fail")
			t.Error(err)
		}

		ret, ok, err := checkEnvironmentVariable(fs, envPath)
		if ret != envPath {
			t.Log("fail")
			t.Fail()
		}
		if !ok {
			t.Log("fail")
			t.Fail()
		}
		if err != nil {
			t.Log("fail")
			t.Error(err)
		}
	})
}

func Test__Config__checkCurrentDirectory(t *testing.T) {
	env := "test"
	defer func() {
		filesystem.ResetTestFs()
	}()

	pattern := []struct {
		Path     string
		EnvPath  string
		Expected bool
	}{
		{"/test/foo/bar", "", false},
		{"/test/foo/bar", "/test/foo/bar/.env", true},
		{"/test/foo/bar", "/test/.env", true},
		{"/test/foo/bar", "/test/.env.test", true},
		{"/test/foo/bar", "/test/.env.production", false},
	}

	for _, example := range pattern {
		filesystem.ResetTestFs()
		fs := filesystem.GetFs()
		fs.MkdirAll(example.Path, 0755)
		if example.EnvPath != "" {
			fs.MkdirAll(filepath.Dir(example.EnvPath), 0755)
			fs.Create(example.EnvPath)
		}

		ret, ok, err := checkCurrentDirectory(fs, env, example.Path)
		if err != nil {
			t.Log("fail")
			t.Error(err)
		}

		if ok != example.Expected {
			t.Log("fail")
			t.Error(err)
		}

		if ok && example.Expected && ret != example.EnvPath {
			t.Logf("fail. example=%v actual=%v", example, ret)
			t.Fail()
		}
	}
}

func Test__Config__checkConfigDirectory(t *testing.T) {
	env := "test"
	tmpHome := os.Getenv("HOME")
	os.Setenv("HOME", "/home/tett23")
	defer func() {
		filesystem.ResetTestFs()
		os.Setenv("HOME", tmpHome)
	}()

	pattern := []struct {
		EnvPath  string
		Expected bool
	}{
		{"/home/tett23/.config/kinsro/.env", true},
		{"/home/tett23/.config/kinsro/.env.test", true},
		{"/home/tett23/.config/kinsro/.env.production", false},
	}

	for _, example := range pattern {
		filesystem.ResetTestFs()
		fs := filesystem.GetFs()
		fs.MkdirAll(filepath.Dir(example.EnvPath), 0755)
		fs.Create(example.EnvPath)

		ret, ok, err := checkConfigDirectory(fs, env, "/home/tett23")
		if err != nil {
			t.Log("fail")
			t.Error(err)
		}

		if ok != example.Expected {
			t.Log("fail")
			t.Error(err)
		}

		if ok && example.Expected && ret != example.EnvPath {
			t.Logf("fail. example=%v actual=%v", example, ret)
			t.Fail()
		}
	}
}
