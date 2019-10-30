package encode

import (
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filelock"
)

// Encode Encode
func Encode(conf *config.Config, fs afero.Fs, info *EncodeInfo) error {

	if err := EncodeTSFile(conf, fs, info); err != nil {
		return errors.Wrapf(err, "encode error. ts=%v", info.RawPath)
	}

	if err := info.Move(syscall.Statfs, fs); err != nil {
		return errors.Wrapf(err, "move error. ts=%v", info.RawPath)
	}

	return nil
}

const tsWidth = "1440"
const tsHeight = "1080"

// EncodeTSFile EncodeTSFile
func EncodeTSFile(conf *config.Config, fs afero.Fs, info *EncodeInfo) error {
	commandOptions := []string{
		conf.FfmpegScriptPath,
		info.RawPath,
		info.TSDest(),
		conf.FfmpegPresetPath,
		tsWidth,
		tsHeight,
	}

	return filelock.Filelock(fs, info.RawPath, encodeCommand(fs, info.RawPath, commandOptions))
}

func encodeCommand(fs afero.Fs, tsPath string, commandOptions []string) func() error {
	return func() error {
		// HACK: テスト用
		afero.WriteFile(fs, strings.ReplaceAll(tsPath, ".ts", ".mp4"), []byte{}, 0644)

		cmd := exec.Command("sh", commandOptions...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		return cmd.Run()
	}
}
