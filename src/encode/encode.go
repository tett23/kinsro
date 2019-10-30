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

	// エンコードだけで終了。moveは別の責務とする

	if err := info.Move(syscall.Statfs, fs); err != nil {
		return errors.Wrapf(err, "move error. ts=%v", tsPath)
	}

	return nil
}

// EncodeTSFile EncodeTSFile
func EncodeTSFile(conf *config.Config, fs afero.Fs, tsPath string) error {
	commandOptions := []string{
		conf.FfmpegScriptPath,
		tsPath,
		conf.FfmpegPresetPath,
		"1440",
		"1080",
	}

	filelock.Filelock(fs, tsPath, func() error {
		// HACK: テスト用
		afero.WriteFile(fs, strings.ReplaceAll(tsPath, ".ts", ".mp4"), []byte{}, 0444)

		// afero.WriteFile(fs, tsPath+"lock", strconv.)
		cmd := exec.Command("sh", commandOptions...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	})

	return nil
}
