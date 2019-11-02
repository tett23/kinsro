package encode

import (
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/fileentry/mpegts"
	"github.com/tett23/kinsro/src/filelock"
)

const tsWidth = "1440"
const tsHeight = "1080"

// Encode Encode
func Encode(conf *config.Config, fs afero.Fs, ts *mpegts.MpegTS) error {
	commandOptions := []string{
		conf.FfmpegScriptPath,
		ts.Src(),
		ts.Dest(),
		conf.FfmpegPresetPath,
		tsWidth,
		tsHeight,
	}

	// HACK: テスト用
	afero.WriteFile(fs, strings.ReplaceAll(ts.Src(), ".ts", ".mp4"), []byte{}, 0644)

	return filelock.Filelock(fs, ts.Src(), encodeCommand(fs, commandOptions))
}

func encodeCommand(fs afero.Fs, commandOptions []string) func() error {
	return func() error {
		cmd := exec.Command("sh", commandOptions...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		return cmd.Run()
	}
}
