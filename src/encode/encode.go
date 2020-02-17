package encode

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
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
	afero.WriteFile(fs, ts.Dest(), []byte{}, 0644)

	return filelock.Filelock(fs, ts.Src(), encodeCommand(fs, ts, commandOptions))
}

func encodeCommand(fs afero.Fs, ts *mpegts.MpegTS, commandOptions []string) func() error {
	return func() error {
		logFile, err := fs.OpenFile(ts.LogPath(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return errors.Wrapf(err, "")
		}

		cmd := exec.Command("sh", commandOptions...)

		cmd.Stdout = logFile
		cmd.Stderr = logFile

		return cmd.Run()
	}
}
