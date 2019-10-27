package encode

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

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

	if err := CreateEncodeLock(fs, tsPath); err != nil {
		return err
	}

	// HACK: テスト用
	afero.WriteFile(fs, strings.ReplaceAll(tsPath, ".ts", ".mp4"), []byte{}, 0444)

	// afero.WriteFile(fs, tsPath+"lock", strconv.)
	cmd := exec.Command("sh", commandOptions...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
	defer Free(fs, tsPath)

	return nil
}

// IsFree IsFree
func IsFree(fs afero.Fs, tsPath string) (bool, error) {
	ok, _ := afero.Exists(fs, tsPath+".lock")
	if !ok {
		return true, nil
	}

	lockTo, err := readLockTimestamp(fs, tsPath)
	if err != nil {
		return false, err
	}

	now := time.Now().Unix()

	return lockTo < now, nil
}

// Free Free
func Free(fs afero.Fs, tsPath string) error {
	ok, _ := afero.Exists(fs, tsPath+".lock")
	if !ok {
		return nil
	}

	return fs.Remove(tsPath + ".lock")
}

func readLockTimestamp(fs afero.Fs, tsPath string) (int64, error) {
	data, err := afero.ReadFile(fs, tsPath+".lock")
	if err != nil {
		return -1, err
	}

	lockTo, err := strconv.Atoi(string(data))
	if err != nil {
		return -1, err
	}

	return int64(lockTo), nil
}

// CreateEncodeLock CreateEncodeLock
func CreateEncodeLock(fs afero.Fs, tsPath string) error {
	ok, err := IsFree(fs, tsPath)
	if err != nil {
		return err
	}
	if !ok {
		lockTo, _ := readLockTimestamp(fs, tsPath)
		return errors.Errorf("file locked. path=%v %v", tsPath, lockTo)
	}

	now := time.Now().Unix()
	lockTo := now + 60*60*24

	timestampString := strconv.Itoa(int(lockTo))

	return afero.WriteFile(fs, tsPath+".lock", []byte(timestampString), 0744)
}
