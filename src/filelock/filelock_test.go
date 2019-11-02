package filelock_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/filelock"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestFilelock__Filelock(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		err := filelock.Filelock(fs, "/test.ts", func() error { return nil })
		assert.NoError(t, err)

		ok, _ := afero.Exists(fs, "/test.ts.lock")
		assert.False(t, ok)
	})

	t.Run("callback failed", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		err := filelock.Filelock(fs, "/test.ts", func() error { return errors.New("") })
		assert.Error(t, err)
	})

	t.Run("lock error", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("2147483647"), 0744)

		err := filelock.Filelock(fs, "/test.ts", func() error { return errors.New("") })
		assert.Error(t, err)
	})
}

func TestEncode__IsFree(t *testing.T) {
	t.Run("lock file does not exists", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		ok, err := filelock.IsFree(fs, "/test.ts")
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("lock file has been expired", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("10"), 0744)

		ok, err := filelock.IsFree(fs, "/test.ts")
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("lock file has not been expired", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("2147483647"), 0744)

		ok, err := filelock.IsFree(fs, "/test.ts")
		assert.NoError(t, err)
		assert.False(t, ok)
	})

	t.Run("when invalid lock", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("invalid lock"), 0744)

		ok, err := filelock.IsFree(fs, "/test.ts")
		assert.Error(t, err)
		assert.False(t, ok)
	})
}

func TestEncode__Lock(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		err := filelock.Lock(fs, "/test.ts")
		assert.NoError(t, err)

		ok, _ := afero.Exists(fs, "/test.ts.lock")
		assert.True(t, ok)
	})

	t.Run("file already locked", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("2147483647"), 0744)

		err := filelock.Lock(fs, "/test.ts")
		assert.NotNil(t, err)
	})

	t.Run("when invalid lock", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("invalid lock"), 0744)

		err := filelock.Lock(fs, "/test.ts")
		assert.NotNil(t, err)
	})
}
