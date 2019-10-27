package encode

import (
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
)

// Encode Encode
func Encode(conf *config.Config, fs afero.Fs, tsPath string) error {
	info, err := NewEncodeInfo(conf, fs, tsPath)
	if err != nil {
		return errors.Wrapf(err, "initialize error. ts=%v", tsPath)
	}

	if err := encodeTS(tsPath); err != nil {
		return errors.Wrapf(err, "encode error. ts=%v", tsPath)
	}

	if err := info.Move(syscall.Statfs, fs); err != nil {
		return errors.Wrapf(err, "move error. ts=%v", tsPath)
	}

	return nil
}

func encodeTS(tsPath string) error {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		println(err.Error())
		return err
	}
	println(string(out))

	return nil
}

// 環境によってshellscriptとpresetのパスをかえる
// デプロイでシェルとpresetを送る
