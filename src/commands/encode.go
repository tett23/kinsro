package commands

import (
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/encode"
	"github.com/tett23/kinsro/src/filesystem"
)

// EncodeTS EncodeTS
func EncodeTS(conf *config.Config, tsPath string) error {
	fs := filesystem.GetFs()
	encode.Encode(conf, fs, tsPath)

	return nil
}