package encode

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/fileentry/mpegts"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestEncode__Encode(t *testing.T) {
	conf := config.GetConfig()

	t.Run("ok", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, "/20190101_foo.ts", []byte{}, 0744)
			ts, _ := mpegts.NewMpegTS("/20190101_foo.ts")

			err := Encode(conf, fs, ts)
			assert.NoError(t, err)
		})
	})
}
