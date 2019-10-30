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
	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		err := filelock.Filelock(fs, "/test.ts", func() error { return nil })
		assert.Nil(t, err)

		ok, _ := afero.Exists(fs, "/test.ts.lock")
		assert.False(t, ok)
	})

	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		err := filelock.Filelock(fs, "/test.ts", func() error { return errors.New("") })
		assert.NotNil(t, err)
	})
}

func TestEncode__IsFree(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		ok, err := filelock.IsFree(fs, "/test.ts")
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("10"), 0744)

		ok, err := filelock.IsFree(fs, "/test.ts")
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("2147483647"), 0744)

		ok, err := filelock.IsFree(fs, "/test.ts")
		assert.Nil(t, err)
		assert.False(t, ok)
	})
}

func TestEncode__Lock(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		err := filelock.Lock(fs, "/test.ts")
		assert.Nil(t, err)

		ok, _ := afero.Exists(fs, "/test.ts.lock")
		assert.True(t, ok)
	})

	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		err := filelock.Lock(fs, "/test.ts")
		assert.Nil(t, err)

		err = filelock.Lock(fs, "/test.ts")
		assert.NotNil(t, err)
	})
}
