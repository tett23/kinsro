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
		afero.WriteFile(fs, "/test/20190101_foo.ts", []byte{}, 0744)

		err := encode.Encode(conf, fs, "/test/20190101_foo.ts")
		assert.Nil(t, err)
	})
}

func TestEncode__EncodeTSFile(t *testing.T) {
	t.Run("", func(t *testing.T) {
		conf := config.GetConfig()
		fs := filesystem.ResetTestFs()
		fs.MkdirAll("/test", 0744)
		afero.WriteFile(fs, "/test/20190101_foo.ts", []byte{}, 0744)

		err := encode.EncodeTSFile(conf, fs, "/test/20190101_foo.ts")
		assert.Nil(t, err)

		ok, _ := afero.Exists(fs, "/test/20190101_foo.ts")
		assert.True(t, ok)

		ok, _ = afero.Exists(fs, "/test/20190101_foo.mp4")
		assert.True(t, ok)

		ok, _ = afero.Exists(fs, "/test/20190101_foo.ts.lock")
		assert.False(t, ok)
	})

	t.Run("TS locked", func(t *testing.T) {
		conf := config.GetConfig()
		fs := filesystem.ResetTestFs()
		fs.MkdirAll("/test", 0744)
		afero.WriteFile(fs, "/test/20190101_foo.ts", []byte{}, 0744)
		afero.WriteFile(fs, "/test/20190101_foo.ts.lock", []byte("2147483647"), 0744)

		err := encode.EncodeTSFile(conf, fs, "/test/20190101_foo.ts")
		assert.NotNil(t, err)
	})
}
