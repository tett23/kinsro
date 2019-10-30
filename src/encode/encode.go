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
func Encode(conf *config.Config, fs afero.Fs, tsPath string) error {
	info, err := NewEncodeInfo(conf, fs, tsPath)
	if err != nil {
		return errors.Wrapf(err, "initialize error. ts=%v", tsPath)
	}

	if err := EncodeTSFile(conf, fs, tsPath); err != nil {
		return errors.Wrapf(err, "encode error. ts=%v", tsPath)
	}

	if err := info.Move(syscall.Statfs, fs); err != nil {
		return errors.Wrapf(err, "move error. ts=%v", tsPath)
	}

	return nil
}

const tsWidth = "1440"
const tsHeight = "1080"

// EncodeTSFile EncodeTSFile
func EncodeTSFile(conf *config.Config, fs afero.Fs, tsPath string) error {
	commandOptions := []string{
		conf.FfmpegScriptPath,
		tsPath,
		conf.FfmpegPresetPath,
		tsWidth,
		tsHeight,
	}

	return filelock.Filelock(fs, tsPath, encodeCommand(fs, tsPath, commandOptions))
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
