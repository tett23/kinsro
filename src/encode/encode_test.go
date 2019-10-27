package encode_test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/encode"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestEncode__Encode(t *testing.T) {
	t.Run("", func(t *testing.T) {
		conf := config.GetConfig()
		fs := filesystem.ResetTestFs()
		fs.MkdirAll("/test", 0744)
		afero.WriteFile(fs, "/test/foo.ts", []byte{}, 0744)

		err := encode.EncodeTSFile(conf, fs, "/test/foo.ts")
		assert.Nil(t, err)

		ok, _ := afero.Exists(fs, "/test/foo.ts")
		assert.Equal(t, ok, true)

		ok, _ = afero.Exists(fs, "/test/foo.mp4")
		assert.Equal(t, ok, true)
	})
}

func TestEncode__EncodeTSFile(t *testing.T) {
	t.Run("", func(t *testing.T) {
		conf := config.GetConfig()
		fs := filesystem.ResetTestFs()
		fs.MkdirAll("/test", 0744)
		afero.WriteFile(fs, "/test/foo.ts", []byte{}, 0744)

		err := encode.EncodeTSFile(conf, fs, "/test/foo.ts")
		assert.Nil(t, err)

		ok, _ := afero.Exists(fs, "/test/foo.ts")
		assert.True(t, ok)

		ok, _ = afero.Exists(fs, "/test/foo.mp4")
		assert.True(t, ok)

		ok, _ = afero.Exists(fs, "/test/foo.ts.lock")
		assert.False(t, ok)
	})

	t.Run("", func(t *testing.T) {
		conf := config.GetConfig()
		fs := filesystem.ResetTestFs()
		fs.MkdirAll("/test", 0744)
		afero.WriteFile(fs, "/test/foo.ts", []byte{}, 0744)
		afero.WriteFile(fs, "/test/foo.ts.lock", []byte("2147483647"), 0744)

		err := encode.EncodeTSFile(conf, fs, "/test/foo.ts")
		assert.NotNil(t, err)
	})
}

func TestEncode__IsLocked(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		ok, err := encode.IsFree(fs, "/test.ts")
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("10"), 0744)

		ok, err := encode.IsFree(fs, "/test.ts")
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, "/test.ts.lock", []byte("2147483647"), 0744)

		ok, err := encode.IsFree(fs, "/test.ts")
		assert.Nil(t, err)
		assert.False(t, ok)
	})
}

func TestEncode__CreateEncodeLock(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fs := filesystem.ResetTestFs()

		err := encode.CreateEncodeLock(fs, "/test.ts")
		assert.Nil(t, err)

		ok, _ := afero.Exists(fs, "/test.ts.lock")
		assert.True(t, ok)
	})
}
